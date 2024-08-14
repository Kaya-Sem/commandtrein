package cmd

// https://docs.irail.be/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FIX: this is appearantly a security issue

// when a time is not specified for timetables, we should use the NotTimed URL for better responses from the API.
const (
	iRailAPIBaseURL = "https://api.irail.be"
	allStationsURL  = iRailAPIBaseURL + "/stations/?format=json&lang=nl"
)

// https://docs.irail.be/#liveboard-liveboard-api-get
func GetSNCBStationTimeTable(stationName string, time string, arrdep string) ([]byte, error) {
	url := fmt.Sprintf("https://api.irail.be/liveboard/?station=%s&lang=nl&format=json", stationName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Parse the iRail departures JSON into a slice of Departure structs
func ParseiRailDepartures(jsonData []byte) ([]timetableDeparture, error) {
	var response StationTimetableResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}

	return response.Departures.Departure, nil
}

func GetSNCBStationsJSON() []byte {
	req, err := http.NewRequest("GET", allStationsURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	return body
}

// Intermediate struct to match the JSON structure
type Station struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	StandardName string `json:"standardname"`
}

var result struct {
	Stations []Station `json:"station"`
}

// https://docs.irail.be/#stations-stations-api-get

func ParseStations(jsonData []byte) ([]Station, error) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v - input data: %s", err, string(jsonData))
	}

	return result.Stations, nil
}
