package cmd

import (
	"time"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/Kaya-Sem/commandtrein/internal/util"
	table "github.com/Kaya-Sem/commandtrein/tables"
	teaTable "github.com/charmbracelet/bubbles/table"
)

func handleConnection(stationFrom string, stationTo string) {
	s := NewSpinner("", " fetching connections", 1*time.Second)
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
		{Title: "Vertrek", Width: 9},
		{Title: "Reistijd", Width: 14},
		{Title: "Aankomst", Width: 8},
		{Title: "Spoor", Width: 10},
	}

	rows := make([]teaTable.Row, len(connections))

	for i, conn := range connections {
		departureTimeWithDelay := util.UnixToHHMM(conn.Departure.Time)
		delay := util.FormatDelay(conn.Departure.Delay)
		if delay != "" {
			departureTimeWithDelay += " " + delay
		}

		rows[i] = teaTable.Row{
			departureTimeWithDelay,
			util.GetDurationInMinutes(conn),
			util.UnixToHHMM(conn.Arrival.Time),
			conn.Departure.Platform,
		}
	}

	s.Stop()
	table.RenderTable(columns, rows, connections)
}
