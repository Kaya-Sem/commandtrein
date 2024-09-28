package main

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	table "github.com/Kaya-Sem/commandtrein/cmd/tables"
	"os"
	"time"

	teaTable "github.com/charmbracelet/bubbles/table"
)

func main() {
	// TODO: allow flags for time and arrdep
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
	s := cmd.NewSpinner("", " fetching connections...", 1*time.Second)
	s.Start()

	connectionsJSON, err := api.GetConnections(stationFrom, stationTo)
	if err != nil {
		panic(err)
	}

	connections, err := api.ParseConnections(connectionsJSON)
	if err != nil {
		panic(err)
	}

	columns := []teaTable.Column{
		{Title: "D", Width: 9},
		{Title: "🕑", Width: 7},
		{Title: "A", Width: 7},
		{Title: "T", Width: 10},
	}

	rows := make([]teaTable.Row, len(connections))

	for i, conn := range connections {
		// Append the formatted delay to the departure time
		departureTimeWithDelay := cmd.UnixToHHMM(conn.Departure.Time)
		delay := cmd.FormatDelay(conn.Departure.Delay)
		if delay != "" {
			departureTimeWithDelay += delay // Append delay to the time
		}

		// Populate the row
		rows[i] = teaTable.Row{
			departureTimeWithDelay,
			api.GetDurationInMinutes(conn),
			cmd.UnixToHHMM(conn.Arrival.Time),
			conn.Departure.Platform,
		}
	}

	s.Stop()
	table.RenderTable(columns, rows, connections)

}

func handleSearch() {
	stationsJSON, err := api.GetSNCBStationsJSON()
	stations, err := api.ParseStations(stationsJSON)
	if err != nil {
		panic(err)
	}

	for _, station := range stations {
		fmt.Printf("%s %s\n", station.ID, station.Name)
	}
}

func handleTimetable(stationName string) {
	s := cmd.NewSpinner("", " fetching timetable...", 1*time.Second)
	s.Start()

	timetableJSON, err := api.GetSNCBStationTimeTable(stationName)
	if err != nil {
		panic(err)
	}

	departures, err := api.ParseiRailDepartures(timetableJSON)
	if err != nil {
		fmt.Printf("failed to parse iRail departures JSON: %v", err)
	}

	columns := []teaTable.Column{
		{Title: "", Width: 5},
		{Title: "", Width: 5},
		{Title: "Destination", Width: 20},
		{Title: "Track", Width: 10},
	}

	rows := make([]teaTable.Row, len(departures))

	for i, departure := range departures {
		var delay string
		if departure.Delay == "0" {
			delay = ""
		} else {
			delay = cmd.FormatDelay(departure.Delay)
		}
		rows[i] = teaTable.Row{
			cmd.UnixToHHMM(departure.Time),
			delay,
			departure.Station,
			departure.Platform,
		}
	}

	s.Stop()

	table.RenderTable(columns, rows, departures)
}
