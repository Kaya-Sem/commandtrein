package main

import (
	"fmt"
	"github.com/Command-Transport/commandtrein/cmd"
	"os"
)

const Version = "0.0.0"

func main() {
	args := cmd.ShiftArgs(os.Args)

	if len(args) == 1 {
		if args[0] == "search" {
			handleSearch()
		} else {
			handleTimetable(args[0])
		}

	} else if len(args) == 3 {
		handleConnection(args[0], args[2])
	}

}

func handleConnection(stationFrom string, stationTo string) {
	connectionsJSON, err := cmd.GetConnections(stationFrom, stationTo, "", "")
	if err != nil {
		panic(err)
	}

	connections, err := cmd.ParseConnections(connectionsJSON)
	if err != nil {
		panic(err)
	}

	// TODO: simple flag for basic lines
	/* 	cmd.PrintConnection(connections) */
	cmd.PrintDepartureTable(connections)
}

func handleSearch() {
	stationsJSON := cmd.GetSNCBStationsJSON()
	stations, err := cmd.ParseStations(stationsJSON)
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Printf("%s %s\n", station.ID, station.Name)
	}
}

func handleTimetable(stationName string) {
	timetableJSON, err := cmd.GetSNCBStationTimeTable(stationName, "", "departure")
	if err != nil {
		panic(err)
	}

	departures, err := cmd.ParseiRailDepartures(timetableJSON)
	if err != nil {
		fmt.Printf("failed to parse iRail departures JSON: %v", err)
	}

	for _, departure := range departures {
		cmd.PrintDeparture(departure)
	}
}
