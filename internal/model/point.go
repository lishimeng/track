package model

type Point struct {
	Time      int64
	Longitude float64
	Latitude  float64
	Speed     float64
}

type TruckHis struct {
	TruckNo string

	Points []Point
}

type TruckSummary struct {
	TruckNo     string
	TotalPoints int
	StartPoint  int
	EndPoint    int
}
