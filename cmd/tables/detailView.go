package table

import (
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
)

func buildDetailView(conn api.Connection) string {
	yellow := "\033[33m"
	italic := "\033[3m" // ANSI escape code for italic
	red := "\033[31m"   // ANSI escape code for red
	reset := "\033[0m"  // Reset color

	// Start building the output string
	output := "\n"
	output += " " + cmd.UnixToHHMM(conn.Departure.Time) + "  " + yellow + "S" + reset + " " + conn.Departure.Station + "\n"
	output += " " + wred + cmd.FormatDelay(conn.Departure.Delay) + reset + yellow + "    ┃" + reset + " " + italic + conn.Departure.VehicleInfo.ShortName + reset + "\n"
	output += yellow + "        ┃ " + reset + "\n"
	output += yellow + "        ┃ " + reset + "\n"

	// Uncomment and modify the following code to add stops
	// for i, stop := range conn.Vias.Via {
	// 	if i == len(conn.Vias.Via)-1 { // Last stop gets blue color
	// 		output += yellow + "      \u2502 " + reset + "\n"
	// 		output += blue + "      ○ " + reset + stop.Station + "\n"
	// 	} else {
	// 		output += yellow + "      \u2502 " + reset + "\n"
	// 		output += yellow + "      ○ " + reset + stop.Station + "\n"
	// 	}
	// }

	return output
}
