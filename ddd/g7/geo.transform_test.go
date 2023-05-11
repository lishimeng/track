package g7

import (
	"github.com/lishimeng/track/internal/model"
	"strconv"
	"testing"
)

func TestGeoTrans(t *testing.T) {
	var lon = "121.51690306558"
	var lat = "31.312694502579"

	latitude, _ := strconv.ParseFloat(lat, 128)
	longitude, _ := strconv.ParseFloat(lon, 128)
	t.Logf("%f, %f\n", longitude, latitude)
}

func TestDistance(t *testing.T) {
	var p1 = model.Point{
		Longitude: 115.9390544,
		Latitude:  25.8830631,
	}
	var p2 = model.Point{
		Longitude: 115.9628511,
		Latitude:  25.8901864,
	}
	var d = distance(p1, p2)
	t.Log(d)
}
