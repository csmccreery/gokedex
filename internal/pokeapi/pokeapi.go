package pokeapi

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationArea struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}

func GetLocations(address string) []LocationArea, err {
	maps := []LocationArea{}
	for address != nil {
		locArr := LocationArea{}
		res, err := http.Get(address)
		if err := nil {
			return nil, err
		}

		res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			return nil, fmt.Errorf("HTTP request to %s failed with status code %s", address, res.StatusCode)
		}

		if err != nil {
			return nil, err
		}

		err := json.Unmarshal(body, &locArr)
		if err != nil {
			return nil, err
		}

		maps = append(maps, locArr)
		address = LocArr.Next
	}

	return maps, nil
}
