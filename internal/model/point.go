package model

// Point
// 轨迹点
type Point struct {
	Time      int64
	Longitude float64
	Latitude  float64
	Speed     float64
}

// TruckHis
// 轨迹历史
type TruckHis struct {
	TruckNo string
	Points  []Point
}

// TruckSummary
// 轨迹汇总
type TruckSummary struct {
	TruckNo     string
	TotalPoints int
	StartPoint  int
	EndPoint    int
}
