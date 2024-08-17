package main

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/charmbracelet/bubbles/table"
)

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
	s.Suffix = " fetching connections..."
	s.Start()
	time.Sleep(1 * time.Second)

	connectionsJSON, err := api.GetConnections(stationFrom, stationTo, "", "")
	if err != nil {
		panic(err)
	}

	connections, err := api.ParseConnections(connectionsJSON)
	if err != nil {
		panic(err)
	}

	s.Stop()
	cmd.PrintDepartureTable(connections)
}

func handleSearch() {
	stationsJSON := api.GetSNCBStationsJSON()
	stations, err := api.ParseStations(stationsJSON)
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Printf("%s %s\n", station.ID, station.Name)
	}
}

func handleTimetable(stationName string) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = "  "
	s.Suffix = " fetching timetable..."
	s.Start()
	time.Sleep(1 * time.Second)

	timetableJSON, err := api.GetSNCBStationTimeTable(stationName, "", "departure")
	if err != nil {
		panic(err)
	}

	departures, err := api.ParseiRailDepartures(timetableJSON)
	if err != nil {
		fmt.Printf("failed to parse iRail departures JSON: %v", err)
	}

	columns := []table.Column{
		{Title: "", Width: 5},
		{Title: "", Width: 4},
		{Title: "Destination", Width: 20},
		{Title: "Track", Width: 10},
	}

	rows := make([]table.Row, len(departures))

	for i, departure := range departures {
		var delay string
		if departure.Delay == "0" {
			delay = ""
		} else {
			delay = cmd.FormatDelay(departure.Delay)
		}
		rows[i] = table.Row{
			cmd.UnixToHHMM(departure.Time),
			delay,
			departure.Station,
			departure.Platform,
		}
	}

	s.Stop()

	cmd.RenderTimetableTable(columns, rows, departures)
}
