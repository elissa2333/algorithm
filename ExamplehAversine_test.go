package algorithm

import "fmt"

func ExampleHaversineFormula() {
	km := HaversineFormula(Point{Lng: 103.949005, Lat: 30.748611}, Point{104.142639, 30.878349})
	fmt.Println(km)
	// Output:
	// 23.45337371103247
}
