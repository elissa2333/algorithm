package algorithm

import (
	"math"
)

// Point 地球上经纬度代表的一个点
type Point struct {
	Lng float64 // 经度
	Lat float64 // 纬度
}

// HaversineFormula 半正矢公式
// https://en.wikipedia.org/wiki/Haversine_formula
// 返回地球上两点之间的距离
// 因为地球不是绝对的圆所以存在误差
func HaversineFormula(p1, p2 Point) (km float64) {
	radius := 6371000.0 //6378137.0 // 地球平均半径
	rad := math.Pi / 180.0
	p1.Lat = p1.Lat * rad
	p1.Lng = p1.Lng * rad

	p2.Lat = p2.Lat * rad
	p2.Lng = p2.Lng * rad

	theta := p2.Lng - p1.Lng
	acos := math.Acos(math.Sin(p1.Lat)*math.Sin(p2.Lat) + math.Cos(p1.Lat)*math.Cos(p2.Lat)*math.Cos(theta))

	return acos * radius / 1000
}
