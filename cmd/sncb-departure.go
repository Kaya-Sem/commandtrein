package cmd

import (
	"fmt"
	"strconv"
)

// PrintDeparture prints the departure information, converting the Unix timestamp to a human-readable format
func PrintDeparture(d timetableDeparture) {
	// Convert Unix timestamp to seconds
	time, err := strconv.ParseInt(d.Time, 10, 64)
	if err != nil {
		fmt.Printf(`Could not convert string %s`, d.Time)
	}
	departureTime := UnixToHHMM(time)
	if err != nil {
		fmt.Printf("Error converting time: %v\n", err)
		return
	}

	// delay is represented as seconds.
	delay, _ := strconv.Atoi(d.Delay)
	delay = delay / 60
	if delay > 0 {
		departureTime = fmt.Sprintf("+%s\033[31m+%s\033[0m", departureTime, FormatDelay(delay))
	}

	fmt.Printf("↳ %s at %s, Platform: %s\n", d.Station, departureTime, d.Platform)
}

type TimetableEntry interface {
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
	Departure []timetableDeparture `json:"departure"`
}

type timetableDeparture struct {
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
