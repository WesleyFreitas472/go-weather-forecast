package services

import (
	"wesleyfreitas472/go-weather-forecast/pkg/utils"
)

type WeatherInterface interface {
	CurrentWeather(lat, long string) (Weather, error)
}

type WeatherImpl struct {
	URL   string
	Token string
}

func (wi WeatherImpl) CurrentWeather(lat, long string) (Weather, error) {
	payload := map[string]string{
		"appid": wi.Token,
		"lat":   lat,
		"lon":   long,
		"units": "imperial",
	}

	client := HttpServiceImpl{}
	endpoint := wi.URL + "/weather"
	resp, err := client.Get(endpoint, payload)
	if err != nil {
		return Weather{}, err
	}
	var weather Weather
	err = utils.UnmarshalPayload(resp, &weather)
	return weather, err
}
