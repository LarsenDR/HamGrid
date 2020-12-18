package main

import (
	"github.com/kv0s/HamGrid/geolocation"
)

func main() {
	var lt geolocation.GeoLocation
	var bt geolocation.BoxLocation

	lt.DLat = 38.8939
	lt.DLong = -92.3598
	// str, err := lt.Grid()
	// if err != nil {
	// 	fmt.Printf("%s/n", err)
	// }
	// fmt.Printf("%v \n", str)

	bt.GridString = "EM38tv"
	bt.Level = 1
	_ = bt.LatLong()
}
