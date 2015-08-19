package main

import (
	"fmt"

	"github.com/scarlson/locate"
	"github.com/scarlson/weather"
)

var (
	// https://developer.forecast.io api key
	apiKey = "YOURKEYHERE"
)

func main() {
	location, err := locate.WhereAmI()
	if err != nil {
		panic(err)
	}
	w := weather.NewEngine(apiKey)
	forecast, err := w.GetForecast(location.Latitude, location.Longitude)
	if err != nil {
		panic(err)
	}
	fmt.Println(forecast.Current.Summary, forecast.Current.Temperature)
}
