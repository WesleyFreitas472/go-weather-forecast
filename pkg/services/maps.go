package services

import (
	"fmt"
	"os"
	"wesleyfreitas472/go-weather-forecast/pkg/utils"
)

type MapsInterface interface {
	GetLocation(address string) (Location, error)
}

type MapsImpl struct {
	URL string
	Key string
}

func (maps MapsImpl) GetLocation(address string) (Location, error) {
	params := map[string]string{
		"address": address,
		"key":     maps.Key,
	}
	url := maps.URL + "/geocode/json"
	client := HttpServiceImpl{}
	resp, err := client.Post(url, params)
	if err != nil {
		return Location{}, nil
	}
	if resp.StatusCode != 200 {
		fmt.Println("Erro ao obter latitude e longitude.")
		os.Exit(0)
	}

	response := MapsResponse{}
	err = utils.UnmarshalPayload(resp, &response)
	if err != nil {
		return Location{}, err
	}
	return response.Results[0].Geometry.Location, nil
}
