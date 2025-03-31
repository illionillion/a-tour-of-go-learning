package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    { Lat: 37.42202, Long: -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Tokyo"] = Vertex{
		Lat: 35.6895, Long: 139.6917,
	}
	for name, vertex := range m {
		fmt.Printf("%s: %+v\n", name, vertex)
	}

	for name, vertex := range m2 {
		fmt.Printf("%s: %+v\n", name, vertex)
	}

}
