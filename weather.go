package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	url = "https://api.forecast.io/forecast/"
)

type Engine struct {
	apiKey string
	url    string
}

func NewEngine(api string) *Engine {
	e := &Engine{}
	e.apiKey = api
	e.url = url + e.apiKey + "/"
	return e
}

func (e *Engine) GetForecast(latitude, longitude float64) (*Forecast, error) {
	lat := strconv.FormatFloat(latitude, 'f', 4, 64)
	long := strconv.FormatFloat(longitude, 'f', 4, 64)
	resp, err := http.Get(e.url + lat + "," + long)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	f := &Forecast{}
	err = json.Unmarshal(bytes, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

type Forecast struct {
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Timezone  string     `json:"timezone"`
	Offset    int        `json:"offset"`
	Current   conditions `json:"currently"`
	Hourly    hourly     `json:"hourly"`
	Daily     hourly     `json:"daily"`
	Units     string     `json:"units"`
}

type hourly struct {
	Summary    string       `json:"summary"`
	Icon       string       `json:"icon"`
	Conditions []conditions `json:"data"`
}

type conditions struct {
	Time                   int     `json:"time"`
	Summary                string  `json:"summary"`
	Icon                   string  `json:"icon"`
	Sunrise                int     `json:"sunriseTime"`
	Sunset                 int     `json:"sunsetTime"`
	MoonPhase              float64 `json:"moonPhase"`
	NearestStormDistance   int     `json:"nearestStormDistance"`
	NearestStormBearing    int     `json:"nearestStormBearing"`
	PrecipIntensity        float64 `json:"precipIntensity"`
	PrecipProbability      float64 `json:"precipProbability"`
	PrecipType             string  `json:"precipType"`
	Temperature            float64 `json:"temperature"`
	TemperatureMin         float64 `json:"temperatureMin"`
	TemperatureMax         float64 `json:"temperatureMax"`
	ApparentTemperature    float64 `json:"apparentTemperature"`
	ApparentTemperatureMin float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMax float64 `json:"apparentTemperatureMax"`
	DewPoint               float64 `json:"dewPoint"`
	Humidity               float64 `json:"humidity"`
	WindSpeed              float64 `json:"windSpeed"`
	WindBearing            int     `json:"windBearing"`
	Visibility             float64 `json:"visibility"`
	CloudCover             float64 `json:"cloudCover"`
	Pressure               float64 `json:"pressure"`
	Ozone                  float64 `json:"ozone"`
}
