package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// makeAPIRequest is a generic function to make HTTP GET requests
func makeAPIRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("couldn't close response body: %v", err))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return body, nil
}

// GetSNCBStationTimeTable fetches the timetable for a specific station
func GetSNCBStationTimeTable(stationName string) ([]byte, error) {
	url := fmt.Sprintf("https://api.irail.be/liveboard/?station=%s&lang=nl&format=json", stationName)
	return makeAPIRequest(url)
}

// ParseiRailDepartures handles fetching of timetable departures
func ParseiRailDepartures(jsonData []byte) ([]TimetableDeparture, error) {
	var response StationTimetableResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}
	return response.Departures.Departure, nil
}

// GetSNCBStationsJSON fetches the list of SNCB stations
func GetSNCBStationsJSON() ([]byte, error) {
	const url string = "https://api.irail.be/stations/?format=json&lang=nl"
	return makeAPIRequest(url)
}

type Station struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StandardName string `json:"standardname"`
}

// ParseStations parses the JSON data of stations
func ParseStations(jsonData []byte) ([]Station, error) {
	var result struct {
		Stations []Station `json:"station"`
	}
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}
	return result.Stations, nil
}
