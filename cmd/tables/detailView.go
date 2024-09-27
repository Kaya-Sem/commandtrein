package table

import (
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
)

func buildDetailView(conn api.Connection) string {
	yellow := "\033[33m"
	italic := "\033[3m" // ANSI escape code for italic
	red := "\033[31m"   // ANSI escape code for red
	reset := "\033[0m"  // ANSI Reset color

	// FIX: better color formatting
	// Start building the output string
	output := "\n"
	output += " " + cmd.UnixToHHMM(conn.Departure.Time) + "  " + yellow + "┏━" + reset + " " + conn.Departure.Station + "\n"
	output += " " + red + cmd.FormatDelay(conn.Departure.Delay) + reset + yellow + "    ┃" + reset + "  " + italic + conn.Departure.VehicleInfo.ShortName + reset + "\n"
	output += yellow + "        ┃ " + reset + "\n"
	output += yellow + "        ┃ " + reset + "\n"

	for _, stop := range conn.Vias.Via {
		output += " " + cmd.UnixToHHMM(conn.Arrival.Time) + yellow + "  ┗━ " + reset + stop.Arrival.Station + ", platform " + stop.Arrival.Platform + "\n"
		output += "        ┊ " + "\n"
		output += " " + cmd.UnixToHHMM(stop.Departure.Time) + yellow + "  ┏━ " + reset + stop.Departure.Station + ", platform " + stop.Departure.Platform + "\n"
		output += yellow + "        ┃ " + reset + "\n"
		output += yellow + "        ┃ " + reset + "\n"
		output += yellow + "        ┃ " + reset + "\n"
	}

	output += " " + cmd.UnixToHHMM(conn.Arrival.Time) + yellow + "  ┗━ " + reset + conn.Arrival.Station + "\n"
	output += "\n"
	return output
}
