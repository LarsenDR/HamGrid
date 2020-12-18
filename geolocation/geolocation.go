package geolocation

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//GeoLocation structure
type GeoLocation struct {
	DLat       float64
	DLong      float64
	Level      int64
	GridString string
}

//BoxLocation structure
type BoxLocation struct {
	GridString string
	Level      int64
	Corner1Lat float64
	Corner1Lon float64
	Corner2Lat float64
	Corner2Lon float64
	Corner3Lat float64
	Corner3Lon float64
	Corner4Lat float64
	Corner4Lon float64
	CenterLat  float64
	CenterLon  float64
}

//LatLong is a method to convert the location to Grid value
func (bt *BoxLocation) LatLong() error {
	var Upper string
	var L2C1Lon, L2C1Lat float64
	Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var Lower string
	Lower = strings.ToLower(Upper)

	fmt.Printf("%d %s\n", bt.Level, bt.GridString)
	// Level One
	LonOff := 180.0 //Degrees
	LatOff := 90.0  //Degrees
	LonInc1 := 20.0 //Degrees
	LatInc1 := 10.0 //Degrees
	Grid1 := string(bt.GridString[0]) + string(bt.GridString[1])
	L1C1Lon := float64(strings.Index(Upper, string(bt.GridString[0])))*LonInc1 - LonOff
	L1C1Lat := float64(strings.Index(Upper, string(bt.GridString[1])))*LatInc1 - LatOff
	fmt.Printf("L1 C1 %v %v, %v\n", Grid1, L1C1Lat, L1C1Lon)
	fmt.Printf("L1 C2 %v %v, %v\n", Grid1, L1C1Lat+10.0, L1C1Lon)
	fmt.Printf("L1 C3 %v %v, %v\n", Grid1, L1C1Lat+10.0, L1C1Lon+20.0)
	fmt.Printf("L1 C4 %v %v, %v\n", Grid1, L1C1Lat, L1C1Lon+20.0)
	fmt.Printf("L1 Ct %v %v, %v\n\n", Grid1, L1C1Lat+5.0, L1C1Lon+10.0)

	//Level Two
	LonInc2 := 2.0 // Degrees
	LatInc2 := 1.0 // Degrees
	Grid2 := string(bt.GridString[0]) + string(bt.GridString[1]) + string(bt.GridString[2]) + string(bt.GridString[3])

	if val, err := strconv.ParseFloat(string(bt.GridString[2]), 64); err == nil {
		L2C1Lon = L1C1Lon + val*LonInc2
	} else {
		fmt.Println(err)
	}

	if val, err := strconv.ParseFloat(string(bt.GridString[3]), 64); err == nil {
		L2C1Lat = L1C1Lat + val*LatInc2
	} else {
		fmt.Println(err)
	}

	fmt.Printf("L2 C1 %v %v, %v\n", Grid2, L2C1Lat, L2C1Lon)
	fmt.Printf("L2 C2 %v %v, %v\n", Grid2, L2C1Lat+LatInc2, L2C1Lon)
	fmt.Printf("L2 C3 %v %v, %v\n", Grid2, L2C1Lat+1.0, L2C1Lon+LonInc2)
	fmt.Printf("L2 C4 %v %v, %v\n", Grid2, L2C1Lat, L2C1Lon+LonInc2)
	fmt.Printf("L2 Ct %v %v, %v\n\n", Grid2, L2C1Lat+(LatInc2/2.0), L2C1Lon+(LatInc2/2.0))

	// Level Three
	LatInc3 := 2.5 / 60.0 //2.5 minutes in decimal degrees
	LonInc3 := 5.0 / 60.0 //5.0 minutes in decimal degrees
	L3C1Lat := L2C1Lat + float64(strings.Index(Lower, string(bt.GridString[5])))*LatInc3
	L3C1Lon := L2C1Lon + float64(strings.Index(Lower, string(bt.GridString[4])))*LonInc3
	Grid3 := bt.GridString
	fmt.Printf("L3 C1 %v %v, %v\n", Grid3, L3C1Lat, L3C1Lon)
	fmt.Printf("L3 C2 %v %v, %v\n", Grid3, L3C1Lat+LatInc3, L3C1Lon)
	fmt.Printf("L3 C3 %v %v, %v\n", Grid3, L3C1Lat+LatInc3, L3C1Lon+LonInc3)
	fmt.Printf("L3 C3 %v %v, %v\n", Grid3, L3C1Lat, L3C1Lon+LonInc3)
	fmt.Printf("L3 Ct %v %v, %v\n", Grid3, L3C1Lat+(LatInc3/2.0), L3C1Lon+(LatInc3/2.0))

	return nil
}

//Grid is a method to convert the location to Grid value
func (lt *GeoLocation) Grid() (string, error) {
	//str := fmt.Sprintf("Change %2.6f, %2.6f to grid level %d.", lat, long, level)
	var Upper, Lower, str, str1, str2, str3 string
	Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lower = strings.ToLower(Upper)
	var glonl2int, glatl2int int64
	var err error

	// Level One
	// For Longitude
	adjlong := lt.DLong + 180.0
	Glong := int(adjlong / 20.0)
	glonl1str := string(Upper[Glong])
	// fmt.Printf("%#v %s\n", Glong, glonl1str)

	// For Latitude
	adjlat := lt.DLat + 90.0
	Glat := int(adjlat / 10.0)
	glatl1str := string(Upper[Glat])
	// fmt.Printf("%#v %#v %s\n", adjlat, Glat, glatl1str)

	str1 = fmt.Sprintf("%s%s", glonl1str, glatl1str)

	// Level Two
	// For Longitude
	nlong := int64(int64(adjlong/2.0) % 10)
	glonl2int = nlong
	// fmt.Printf("%#v\n", glonl2int)

	// For Latitude
	nlat := int64(adjlat) % 10
	glatl2int = nlat
	// fmt.Printf("%#v %#v\n", adjlat, glatl2int)

	str2 = fmt.Sprintf("%s%s%d%d", glonl1str, glatl1str, glonl2int, glatl2int)

	// Level Three
	// For Longitude
	rlong := (adjlong - 2.0*math.Trunc(adjlong/2.0)) * 60.0
	glong := math.Trunc(rlong / 5.0)
	glonl3str := string(Lower[int(glong)])
	// fmt.Printf("%#v %#v %#v %s\n", adjlong, rlong, glong, glonl3str)

	// For Latitude
	rlat := (adjlat - math.Trunc(adjlat)) * 60.0
	glat := math.Trunc(rlat / 2.5)
	glatl3str := string(Lower[int(glat)])
	// fmt.Printf("%#v %#v %#v %s\n", adjlat, rlat, glat, glatl3str)

	str3 = fmt.Sprintf("%s%s%d%d%s%s", glonl1str, glatl1str, glonl2int, glatl2int, glonl3str, glatl3str)

	// Which level requested

	if lt.Level == 1 {
		str = str1
	} else if lt.Level == 2 {
		str = str2
	} else if lt.Level == 3 {
		str = str3
	} else {
		err = fmt.Errorf("Unknown Level")
	}
	return str, err
}

// TradtoDDeg Convert degrees, minutes and seconds to Decimal Degrees
func TradtoDDeg(deg, min, sec float64, hemi string) (float64, error) {
	var err error
	var sign float64
	if hemi == "N" {
		sign = 1.0
	} else if hemi == "S" {
		sign = -1.0
	} else if hemi == "E" {
		sign = 1.0
	} else if hemi == "W" {
		sign = -1.0
	} else {
		err = fmt.Errorf("%s", "The hemi varable must be N, S, E, W")
	}
	DDeg := math.Copysign((deg + (min / 60) + (sec / 3600)), sign)
	return DDeg, err
}

// DDegtoTrad Convert Decimal Degrees to traditional degrees, minutes and seconds
func DDegtoTrad(ddeg float64, latorlon string) (float64, float64, float64, string, error) {
	var deg, min, sec float64
	var hemi string
	var err error
	if (math.Signbit(ddeg)) && (latorlon == "lat") {
		hemi = "S"
	} else if (math.Signbit(ddeg)) && bool(latorlon == "lon") {
		hemi = "W"
	} else if !(math.Signbit(ddeg)) && bool(latorlon == "lat") {
		hemi = "N"
	} else if !(math.Signbit(ddeg)) && bool(latorlon == "lon") {
		hemi = "E"
	} else {
		err = fmt.Errorf("%s", "The latorlon variable must be lat or lon")
	}
	pddeg := math.Copysign(ddeg, 1.0)
	deg = math.Trunc(pddeg)
	min = math.Trunc((pddeg - deg) * 60)
	sec = (pddeg - deg - min/60) * 3600

	return deg, min, sec, hemi, err
}
