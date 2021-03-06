package main

type Rootobject struct {
	Coord   Coord      `json:"coord"`
	Weather []*Weather `json:"weather"`
	Base    string     `json:"base"`
	Main    Main       `json:"main"`
	Wind    Wind       `json:"wind"`
	Clouds  Clouds     `json:"clouds"`
	Dt      int        `json:"dt"`
	Sys     Sys        `json:"sys"`
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Cod     int        `json:"cod"`
}
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}
type Main struct {
	Temp      float64 `json:"temp"`
	Pressure  float64 `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	SeaLevel  float64 `json:"sea_level"`
	GrndLevel float64 `json:"grnd_level"`
}
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}
type Clouds struct {
	All int `json:"all"`
}
type Sys struct {
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
