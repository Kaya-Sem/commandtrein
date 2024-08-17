package api

// https://docs.irail.be/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://docs.irail.be/#liveboard-liveboard-api-get

func GetSNCBStationTimeTable(stationName string, time string, arrdep string) ([]byte, error) {
	url := fmt.Sprintf("https://api.irail.be/liveboard/?station=%s&lang=nl&format=json", stationName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("couldn't close response body: %v", err))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
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

func GetSNCBStationsJSON() []byte {
	const url string = "https://api.irail.be/stations/?format=json&lang=nl"
	req, err := http.NewRequest("GET", url, nil)
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	return body
}

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
