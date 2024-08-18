package api

import (
	"fmt"
	"strconv"
)

func (d TimetableDeparture) GetUnixDepartureTime() int {
	time, err := strconv.Atoi(d.Time)
	if err != nil {
		fmt.Println("Error converting departure time: %s", d.Time)
	}

	return time
}

func (d TimetableDeparture) GetDelayInSeconds() int {
	delay, err := strconv.Atoi(d.Delay)
	if err != nil {
		fmt.Println("Error converting delay: %s", d.Delay)
	}
	return delay
}

type StationTimetableResponse struct {
	Version     string      `json:"version"`
	Timestamp   string      `json:"timestamp"`
	Station     string      `json:"station"`
	StationInfo StationInfo `json:"stationinfo"`
	Departures  Departures  `json:"departures"`
}

type StationInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	LocationX    string `json:"locationX"`
	LocationY    string `json:"locationY"`
	StandardName string `json:"standardname"`
}

type Departures struct {
	Number    string               `json:"number"`
	Departure []TimetableDeparture `json:"departure"`
}

type TimetableDeparture struct {
	ID                  string       `json:"id"`
	Station             string       `json:"station"`
	StationInfo         StationInfo  `json:"stationinfo"`
	Time                string       `json:"time"`
	Delay               string       `json:"delay"` // seconds
	Canceled            string       `json:"canceled"`
	Left                string       `json:"left"`
	IsExtra             string       `json:"isExtra"`
	Vehicle             string       `json:"vehicle"`
	VehicleInfo         VehicleInfo  `json:"vehicleinfo"`
	Platform            string       `json:"platform"`
	PlatformInfo        PlatformInfo `json:"platforminfo"`
	Occupancy           Occupancy    `json:"occupancy"`
	DepartureConnection string       `json:"departureConnection"`
}

type VehicleInfo struct {
	Name      string `json:"name"`
	ShortName string `json:"shortname"`
	Number    string `json:"number"`
	Type      string `json:"type"`
	LocationX string `json:"locationX"`
	LocationY string `json:"locationY"`
	ID        string `json:"@id"`
}

type PlatformInfo struct {
	Name   string `json:"name"`
	Normal string `json:"normal"`
}

type Occupancy struct {
	ID   string `json:"@id"`
	Name string `json:"name"`
}
