package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Kaya-Sem/commandtrein/internal/util"
)

/*
GetConnections fetches the connection data from the API and returns the response body as a byte slice.

https://docs.irail.be/#connections
*/
func GetConnections(stationFrom string, stationTo string, time string, departure bool) ([]byte, error) {

	timesel := "arrival"
	if departure {
		timesel = "departure"
	}

	new_time := ""
	if time == "" {
		new_time = util.GetBelgiumTimeHHMM()
	} else {
		new_time = time
	}

	url := fmt.Sprintf("https://api.irail.be/connections/?from=%s&to=%s&time=%s&timesel=%s&format=json&lang=nl&typeOfTransport=automatic&alerts=false&results=10",
		stationFrom,
		stationTo,
		new_time,
		timesel,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("time probably not correct: %v", err))
		}
	}(resp.Body)

	// Check HTTP response status code
	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("server error: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

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
	Vias      Vias                `json:"vias,omitempty"`
}

// TODO: implement commented bits if asked for
type ConnectionDeparture struct {
	Station  string `json:"station"`
	Time     string `json:"time"`  // Unix since epoch
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
