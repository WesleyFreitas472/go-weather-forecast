package main

import (
	"fmt"
	"os"
	"strings"
	configs "wesleyfreitas472/go-weather-forecast/pkg/config"
	"wesleyfreitas472/go-weather-forecast/pkg/services"
)

type AppContext struct {
	Wheater services.WeatherInterface
	Maps    services.MapsInterface
	Config  *configs.Configuration
}

const configFile = "app"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: wf \"ADDRESS\"")
		return
	}
	address := os.Args[1]
	address = replaceWhiteSpaces(address)

	appContext := AppContext{}
	config, err := configs.ReadConfig(configFile)
	checkError(err)
	appContext.Config = config

	appContext.Wheater = services.WeatherImpl{
		URL:   appContext.Config.App.Weather.URL,
		Token: appContext.Config.App.Weather.Token,
	}

	appContext.Maps = services.MapsImpl{
		URL: appContext.Config.App.Maps.URL,
		Key: appContext.Config.App.Maps.Key,
	}

	location, err := appContext.Maps.GetLocation(address)
	checkError(err)

	lat := fmt.Sprintf("%f", location.Lat)
	lng := fmt.Sprintf("%f", location.Lng)

	weather, err := appContext.Wheater.CurrentWeather(lat, lng)
	checkError(err)

	fmt.Println("Local: ", weather.Name)
	fmt.Println("Temperatura:", (weather.Main.Temp-32)/1.8, "Cº")
}

func replaceWhiteSpaces(address string) string {
	return strings.Replace(address, " ", "%20", -1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
