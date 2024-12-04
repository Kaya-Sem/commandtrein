package cmd

import (
	"fmt"
	"time"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/Kaya-Sem/commandtrein/internal/util"
	table "github.com/Kaya-Sem/commandtrein/tables"
	teaTable "github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewTimetableCommand())
}

func NewTimetableCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "timetable [station]",
		Short: "Show a timetable for a station",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handleTimetable(args[0])
		},
	}
}

func handleTimetable(stationName string) {
	s := NewSpinner("", " fetching timetable...", 1*time.Second)
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
		{Title: "", Width: 8},
		{Title: "Track", Width: 5},
		{Title: "Destination", Width: 28},
	}

	rows := make([]teaTable.Row, len(departures))

	for i, departure := range departures {
		var delay string
		if departure.Delay == "0" {
			delay = ""
		} else {
			delay = util.FormatDelay(departure.Delay)
		}

		rows[i] = teaTable.Row{
			util.UnixToHHMM(departure.Time) + " " + delay,
			table.LeftPad(departure.Platform, 5),
			departure.Station,
		}
	}

	s.Stop()

	table.RenderTable(columns, rows, departures)
}
