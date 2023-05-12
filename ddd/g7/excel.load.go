package g7

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/track/internal/model"
	"github.com/lishimeng/track/internal/task"
	"github.com/xuri/excelize/v2"
	"math"
	"strconv"
)

func init() {
	task.Register(new(DataHandler))
}

type DataHandler struct {
	C   chan Request
	ctx context.Context
}

// Req
// 任务入口
func (d *DataHandler) Req(r Request) {
	if d.ctx == nil {
		return
	}
	select {
	case <-d.ctx.Done():
	case d.C <- r:
	}
}

func (d *DataHandler) Run(ctx context.Context) {

	d.ctx = ctx
	for {
		select {
		case <-d.ctx.Done():
			return
		case r := <-d.C:
			err := Do(r)
			if err != nil {
				log.Info(err)
			}
		}
	}
}

func Do(r Request) (err error) {
	name := r.DataFile
	filePath := model.FileDir + "/" + name
	his, err := LoadExcelG7(filePath, r.DataSheet)
	if err != nil {
		log.Info(err)
		return
	}
	s, err := summary(his, r.StartPoint, r.EndPoint, float64(r.Threshold))
	if err != nil {
		log.Info(err)
		return
	}
	log.Info("-------------------------------------")
	log.Info("车牌:%s", s.TruckNo)
	log.Info("轨迹点数:%d", s.TotalPoints)
	log.Info("开始点:%d, 结束点:%d", s.StartPoint, s.EndPoint)
	return
}

func LoadExcelG7(name string, sheet string) (truckHis model.TruckHis, err error) {
	file, err := excelize.OpenFile(name)
	if err != nil {
		return
	}

	rows, err := file.GetRows(sheet)
	if err != nil {
		return
	}

	for index, row := range rows {
		if index == 0 {
			continue // title行
		}
		if len(row) < 6 {
			continue
		}

		if index == 1 {
			truckHis.TruckNo = row[0]
		}
		var p model.Point
		p, err = transformPoint(row)
		if err != nil {
			continue
		}
		truckHis.Points = append(truckHis.Points, p)
	}
	return
}

func transformPoint(row []string) (p model.Point, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
	}()
	//timeStr := row[1]  // format eg. 05-10 00:01:06
	//speedStr := row[2] // 速度Km/h
	lonStr := row[3] // 经度
	latStr := row[4] // 纬度
	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return
	}
	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return
	}
	p.Latitude = latitude
	p.Longitude = longitude
	return
}

const earthRadius = 6378.137

const (
	StateStart = iota
	StateFindStart
	StateFindEnd
	StateDone
)

func summary(his model.TruckHis, expectStart model.Point, expectEnd model.Point, threshold float64) (s model.TruckSummary, err error) {
	var state = StateStart
	s.TruckNo = his.TruckNo
	s.TotalPoints = len(his.Points)
	for index, p := range his.Points {
		switch state {
		case StateStart:
			state++
		case StateFindStart:
			var d = distance(expectStart, p)
			if d <= threshold {
				// 找到了
				state++
				s.StartPoint = index
				log.Info("find start point %d", index)
			}
		case StateFindEnd:
			var d = distance(expectEnd, p)
			if d <= threshold {
				// 找到了
				state++
				s.EndPoint = index
				log.Info("find end point %d", index)
			}
		case StateDone:
			log.Info("summary done")
			return
		}

	}
	// 运行到这里说明没找到
	if state == StateFindStart {
		err = errors.New("no start point")
	}
	if state == StateFindEnd {
		err = errors.New("no end point")
	}

	return
}

// double s = 2 * Math.asin(Math.sqrt( +
// Math.cos(radLat1)*Math.cos(radLat2)*Math.pow(Math.sin(b/2),2)));
func distance(p1 model.Point, p2 model.Point) (d float64) {
	var radLat1 = radius(p1.Latitude)
	var radLat2 = radius(p2.Latitude)
	var a = radius(p1.Latitude) - radius(p2.Latitude)
	var b = radius(p1.Longitude) - radius(p2.Longitude)
	var s = 2 * math.Asin(math.Sqrt(math.Pow(math.Sin(a/2), 2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Pow(math.Sin(b/2), 2)))
	d = s * earthRadius
	return
}

func radius(d float64) (r float64) {
	r = d * math.Pi / 180.0
	return
}
