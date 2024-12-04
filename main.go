package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	table "github.com/Kaya-Sem/commandtrein/cmd/tables"
	teaTable "github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

const Version = "1.1.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:   "commandtrein [from] [to]",
		Short: "Command line tool for train schedules and connections",
		Long: `Commandtrein allows you to search for train stations, 
find connections between stations, and view timetables for specific stations.`,
		Args: cobra.MaximumNArgs(2), // Allow up to 2 positional arguments
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 2:
				// Default to "connection" behavior when two arguments are provided
				handleConnection(args[0], args[1])
			default:
				// Show help if no valid default operation can be performed
				fmt.Println("Error: Two arguments required for default behavior (from and to).")
				_ = cmd.Help()
			}
		},
	}

	// Add other commands to the root command
	rootCmd.AddCommand(
		newSearchCommand(),
		newTimetableCommand(),
		newVersionCommand(),
	)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newSearchCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "search",
		Short: "Search for train stations",
		Run: func(cmd *cobra.Command, args []string) {
			handleSearch()
		},
	}
}

func newTimetableCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "timetable [station]",
		Short: "Show a timetable for a station",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			handleTimetable(args[0])
		},
	}
}

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the version of commandtrein",
		Run: func(cmd *cobra.Command, args []string) {
			handleVersion()
		},
	}
}

func handleConnection(stationFrom string, stationTo string) {
	s := cmd.NewSpinner("", " fetching connections", 1*time.Second)
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
		departureTimeWithDelay := cmd.UnixToHHMM(conn.Departure.Time)
		delay := cmd.FormatDelay(conn.Departure.Delay)
		if delay != "" {
			departureTimeWithDelay += " " + delay
		}

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
		fmt.Printf("%s\n", station.Name)
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
			delay = cmd.FormatDelay(departure.Delay)
		}

		rows[i] = teaTable.Row{
			cmd.UnixToHHMM(departure.Time) + " " + delay,
			table.LeftPad(departure.Platform, 5),
			departure.Station,
		}
	}

	s.Stop()

	table.RenderTable(columns, rows, departures)
}

func handleVersion() {
	fmt.Printf("commandtrein %s\n", Version)
}
