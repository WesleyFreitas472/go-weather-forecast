package services

type Weather struct {
	Main Main
	Name string
}

type Main struct {
	Temp float32
}

type Coordinate struct {
	Lat  string
	Long string
}

type MapsResponse struct {
	Results []Result
}

type Result struct {
	Geometry Geometry
}

type Geometry struct {
	Location Location
}

type Location struct {
	Lat float32
	Lng float32
}
