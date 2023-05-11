package process

import (
	"fmt"
	"github.com/kataras/iris/v12/x/errors"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/track/internal/model"
	"github.com/xuri/excelize/v2"
	"strconv"
)

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
		log.Info("index: %d", index)
		if len(row) < 6 {
			continue
		}

		if index == 0 {
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
