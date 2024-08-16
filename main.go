package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Command-Transport/commandtrein/cmd"
	"github.com/briandowns/spinner"
	"github.com/charmbracelet/bubbles/table"
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
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "  "
	s.Suffix = " Getting connections..."
	s.Start()
	time.Sleep(2 * time.Second)
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
	s.Stop()
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

	columns := []table.Column{
		{Title: "", Width: 5},
		{Title: "Destination", Width: 20},
		{Title: "Track", Width: 10},
	}

	var rows []table.Row
	for _, departure := range departures {
		row := table.Row{cmd.UnixToHHMM(departure.Time), departure.Station, departure.Platform}
		rows = append(rows, row)
	}

	cmd.RenderTable(columns, rows)
}
