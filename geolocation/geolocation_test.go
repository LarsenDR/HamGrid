package geolocation

import (
	"fmt"
	"log"
	"testing"
)

func TestGridtoBox(t *testing.T) {
	tests := []struct {
		Name   string
		Grid   string
		Level  int64
		ExpLat float64
		ExpLon float64
	}{
		{"Munich", "JN58td", 3, 48.14666, 11.60833},
		{"Montevideo", "GF15vc", 3, -34.91, -56.21166},
		{"Washington, DC", "FM18lw", 3, 38.92, -77.065},
		{"Wellington", "RE78ir", 3, -41.28333, 174.745},
		{"Newington, CT (W1AW)", "FN31pr", 3, 41.714775, -72.727260},
		{"Palo Alto (K6WRU)", "CM87wj", 3, 37.413708, -122.1073236},
		{"Chattanooga (KI6CQ/4)", "EM75kb", 3, 35.0542, -85.1142},
		{"Columbia (KV0S)", "EM38tv", 3, 38.893944, -92.359832},
	}
	var bt BoxLocation

	for _, tt := range tests {
		tt := tt
		bt.GridString = tt.Grid
		bt.Level = tt.Level
		err := bt.LatLong()
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf(" %v, %v, %v, ExpLat: %v, Predicted: %v\n", tt.Name, bt.GridString, bt.Level, tt.ExpLat, bt.Corner1Lat)
	}
	fmt.Println()
}

func TestTlatLontoGrid(t *testing.T) {
	tests := []struct {
		Name     string
		Lat      float64
		Long     float64
		Level    int64
		Expected string
	}{
		{"Munich", 48.14666, 11.60833, 1, "JN"},
		{"Montevideo", -34.91, -56.21166, 1, "GF"},
		{"Washington, DC", 38.92, -77.065, 1, "FM"},
		{"Wellington", -41.28333, 174.745, 1, "RE"},
		{"Newington, CT (W1AW)", 41.714775, -72.727260, 1, "FN"},
		{"Palo Alto (K6WRU)", 37.413708, -122.1073236, 1, "CM"},
		{"Chattanooga (KI6CQ/4)", 35.0542, -85.1142, 1, "EM"},
		{"Columbia (KV0S)", 38.893944, -92.359832, 1, "EM"},
		{"Columbia (KV0S)", 38.8939, -92.3598, 1, "EM"},
		{"Munich", 48.14666, 11.60833, 2, "JN58"},
		{"Montevideo", -34.91, -56.21166, 2, "GF15"},
		{"Washington, DC", 38.92, -77.065, 2, "FM18"},
		{"Wellington", -41.28333, 174.745, 2, "RE78"},
		{"Newington, CT (W1AW)", 41.714775, -72.727260, 2, "FN31"},
		{"Palo Alto (K6WRU)", 37.413708, -122.1073236, 2, "CM87"},
		{"Chattanooga (KI6CQ/4)", 35.0542, -85.1142, 2, "EM75"},
		{"Columbia (KV0S)", 38.893944, -92.359832, 2, "EM38"},
		{"Munich", 48.14666, 11.60833, 3, "JN58td"},
		{"Montevideo", -34.91, -56.21166, 3, "GF15vc"},
		{"Washington, DC", 38.92, -77.065, 3, "FM18lw"},
		{"Wellington", -41.28333, 174.745, 3, "RE78ir"},
		{"Newington, CT (W1AW)", 41.714775, -72.727260, 3, "FN31pr"},
		{"Palo Alto (K6WRU)", 37.413708, -122.1073236, 3, "CM87wj"},
		{"Chattanooga (KI6CQ/4)", 35.0542, -85.1142, 3, "EM75kb"},
		{"Columbia (KV0S)", 38.893944, -92.359832, 3, "EM38tv"},
		{"Columbia (KV0S)", 38.89, -92.35, 3, "EM38tv"},
		{"Columbia (KV0S)", 38.8939, -92.3598, 4, "EM38tv"},
	}
	var lt GeoLocation

	for _, tt := range tests {
		tt := tt
		lt.DLat = tt.Lat
		lt.DLong = tt.Long
		lt.Level = tt.Level
		str, err := lt.Grid()
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		if tt.Expected == str {
			fmt.Printf(" Good %v, %v, %v, Expected: %v, Predicted: %v\n", tt.Name, lt.DLat, lt.DLong, tt.Expected, str)
		} else {
			fmt.Printf("Wrong %v, %v, %v, Expected: %v, Predicted: %v\n", tt.Name, lt.DLat, lt.DLong, tt.Expected, str)
		}
	}
	fmt.Println()
}

func TestTradtoDDeg(t *testing.T) {
	tests := []struct {
		Name string
		Deg  float64
		Min  float64
		Sec  float64
		Hemi string
		DDeg float64
	}{
		{"Munich", 48.0, 22.0, 45.33, "W", 48.379258},
		{"Montevideo", 34, 56, 34.91, "E", 34.943031},
		{"Test", 30, 15, 50.0, "N", 30.263888889},
	}
	for _, tt := range tests {
		str, err := TradtoDDeg(tt.Deg, tt.Min, tt.Sec, tt.Hemi)
		if err != nil {
			log.Fatal(err)
		}
		if tt.DDeg != str {
			fmt.Printf("Trad %v %v %v %s Ddeg %8.6f %8.6f\n", tt.Deg, tt.Min, tt.Sec, tt.Hemi, str, tt.DDeg)
		}
	}
	fmt.Println()

}

// func TestTradtoDeg(t *testing.T) {
func TestDDegtoTrad(t *testing.T) {
	tests := []struct {
		DDeg     float64
		Latorlon string
		Deg      int64
		Min      int64
		Sec      float64
		Hemi     string
	}{
		{48.379258, "lon", 48, 22, 45.33, "W"},
		{34.943031, "lon", 34, 56, 34.91, "E"},
		{48.14666, "lon", 48, 19, 24.37, "E"},
		{-34.91, "lon", 34, 25, 34.7, "W"},
		{-77.065345, "lon", 77, 4, 53.43, "W"},
		{-34.91, "lat", 34, 25, 34.7, "S"},
		{30.263888889, "lat", 30, 15, 50.0, "N"},
	}
	for _, tt := range tests {
		deg, min, sec, hemi, err := DDegtoTrad(tt.DDeg, tt.Latorlon)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Ddeg %8.6f %s Trad %v %v %v %s \n", tt.DDeg, tt.Latorlon, deg, min, sec, hemi)

	}
	fmt.Println()
}
