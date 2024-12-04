package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// GetConnections fetches the connection data from the API and returns the response body as a byte slice.
func GetConnections(stationFrom string, stationTo string) ([]byte, error) {
	url := fmt.Sprintf("https://api.irail.be/connections/?from=%s&to=%s&timesel=departure&format=json&lang=nl&typeOfTransport=automatic&alerts=false&results=6", stationFrom, stationTo)

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
		fmt.Println(fmt.Errorf("couldn't read response body: %v", err))
	}

	return body, nil
}

// ParseConnections takes the response body as a byte slice and returns an array of Connection structs.
func ParseConnections(body []byte) ([]Connection, error) {
	var result ConnectionResult
	err := json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Connection, nil
}

func GetDurationInMinutes(c Connection) string {
	duration, err := strconv.Atoi(c.Duration)
	if err != nil {
		fmt.Printf("Duration could not be parsed: %d", duration)
	}

	duration /= 60

	return strconv.Itoa(duration) + "m"
}

func (c Connection) GetDelayInSeconds() int {
	delay, err := strconv.Atoi(c.Departure.Delay)

	if err != nil {
		fmt.Printf("Error converting delay: %s", c.Departure.Time)
		os.Exit(1)
	}
	return delay
}

func (c Connection) GetUnixDepartureTime() int {
	depTime, err := strconv.Atoi(c.Departure.Time)
	if err != nil {
		fmt.Printf("Error converting departuretime: %s\n", c.Departure.Time)
	}
	return depTime
}

type ConnectionResult struct {
	Connection []Connection `json:"connection"`
}

type Connection struct {
	ID        string              `json:"id"`
	Departure ConnectionDeparture `json:"departure"`
	Arrival   ConnectionArrival   `json:"arrival"`
	Duration  string              `json:"duration"`
	Number    string              `json:"number"`
	Vias      Vias                `json:"vias,omitempty"` // waarom de * ?
}

// TODO:
type ConnectionDeparture struct {
	Station  string `json:"station"`
	Time     string `json:"time"`  // Unix
	Delay    string `json:"delay"` // seconds
	Canceled string `json:"canceled"`
	Left     string `json:"left"`
	IsExtra  string `json:"isExtra"`
	Vehicle  string `json:"vehicle"`
	Platform string `json:"platform"`
	//Stops    []Stop `json:"stops"`
	VehicleInfo VehicleInfo `json:"vehicleinfo"`
	//	StationInfo  StationInfo  `json:"stationinfo"`
	//  PlatformInfo PlatformInfo `json:"platforminfo"`
}

type ConnectionArrival struct {
	Station      string       `json:"station"`
	StationInfo  StationInfo  `json:"stationinfo"`
	Time         string       `json:"time"`  // Unix
	Delay        string       `json:"delay"` // seconds
	Canceled     string       `json:"canceled"`
	Left         string       `json:"left"`
	Platform     string       `json:"platform"`
	PlatformInfo PlatformInfo `json:"platforminfo"`
}

type Stop struct {
	Station string `json:"station"`
	Time    string `json:"time"`  // Unix
	Delay   string `json:"delay"` // seconds
}

type Vias struct {
	Number string    `json:"number"`
	Via    []ViaInfo `json:"via"`
}

type ViaInfo struct {
	ID        string              `json:"id"`
	Arrival   ConnectionArrival   `json:"arrival"`
	Departure ConnectionDeparture `json:"departure"`
	// Station     string              `json:"station"`
	TimeBetween string `json:"timeBetween"`
	//Vehicle     string              `json:"vehicle"`
}
