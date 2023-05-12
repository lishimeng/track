package g7

import "github.com/lishimeng/track/internal/model"

// Request
// 请求
type Request struct {
	StartPoint model.Point // 开始点
	EndPoint   model.Point // 结束点
	Date       int64       // 日期
	DateStr    string      // 格式化日期yyyy-mm-dd(GPS数据只有月日)
	TruckNo    string      // 车牌号
	DataFile   string      // 数据文件名
	DataSheet  string      // excel sheet name
	Threshold  int         // 偏移阈值(KM)
}
