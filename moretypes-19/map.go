package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Tokyo"] = Vertex{
		35.6895, 139.6917,
	}
	for name, vertex := range m {
		fmt.Printf("%s: %+v\n", name, vertex)
	}
}
