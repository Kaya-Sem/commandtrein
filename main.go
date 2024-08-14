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
			stationsJSON := cmd.GetSNCBStationsJSON()
			stations, err := cmd.ParseStations(stationsJSON)
			if err != nil {
				panic(err)
			}

			for _, station := range stations {
				fmt.Printf("%s %s\n", station.ID, station.Name)
			}
		} else {
			handleTimetable(args[0])
		}

	} else if len(args) == 3 {
		stationFrom := args[0]
		stationTo := args[2]

		connectionsJSON, err := cmd.GetConnections(stationFrom, stationTo, "", "")
		if err != nil {
			panic(err)
		}

		connections, err := cmd.ParseConnections(connectionsJSON)
		if err != nil {
			panic(err)
		}

		for _, conn := range connections {
			cmd.PrintConnection(conn)
		}

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
