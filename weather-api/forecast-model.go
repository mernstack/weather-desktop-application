package weather_api

import "time"

type WeatherData struct {
	Query QueryData `json:"query"`
}

type QueryData struct {
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"created"`
	Lang      string    `json:"lang"`
	Results   Result    `json:"results"`
}

type Result struct {
	Channel ChannelData `json:"channel"`
}

type ChannelData struct {
	Title         string       `json:"title"`
	LastBuildDate string       `json:"lastBuildDate"`
	Location      LocationData `json:"location"`
	Item ItemData `json:"item"`
}

type ItemData struct {
	Condition Condition `json:"condition"`
	Forecast []Forecast `json:"forecast"`
}

type Condition struct {
	Code string `json:"code"`
	Date string `json:"date"`
	Temp string `json:"temp"`
	Text string `json:"text"`
}
type LocationData struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Region  string `json:"region"`
}

type Forecast struct {
	Date string `json:"date"`
	Day  string `json:"day"`
	High int    `json:"high"`
	Low  int    `json:"low"`
	Text string `json:"text"`
}
