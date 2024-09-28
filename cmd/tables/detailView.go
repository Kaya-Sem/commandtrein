package table

import (
	"strings"

	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
)

const (
	yellowCode = "\033[33m"
	redCode    = "\033[31m"
	italicCode = "\033[3m"
	resetCode  = "\033[0m"
)

func yellow(text string) string {
	return yellowCode + text + resetCode
}

func red(text string) string {
	return redCode + text + resetCode
}

func italic(text string) string {
	return italicCode + text + resetCode
}

func addVerticalBar(s string, repetitions int) string {
	verticalBar := yellow("        ┃\n")
	return s + strings.Repeat(verticalBar, repetitions)
}

func addVia(text string, v api.ViaInfo) string {
	text += " " + cmd.UnixToHHMM(v.Arrival.Time) + yellow("  ┗━ ") + v.Arrival.Station + ", platform " + v.Arrival.Platform + "\n"
	text += "        ┊ " + "\n"
	text += " " + cmd.UnixToHHMM(v.Departure.Time) + yellow("  ┏━ ") + v.Departure.Station + ", platform " + v.Departure.Platform + "\n"
	text = addVerticalBar(text, 3)

	return text
}

func addArrivalStation(a api.ConnectionArrival) string {
	return " " + cmd.UnixToHHMM(a.Time) + yellow("  ┗━ ") + a.Station + "\n"
}

func addDepartureStation(d api.ConnectionDeparture) string {
	hi := " " + cmd.UnixToHHMM(d.Time) + yellow("  ┏━ ") + d.Station + "\n"
	hi += " " + red(cmd.FormatDelay(d.Delay)) + yellow("    ┃  ") + italic(d.VehicleInfo.ShortName) + "\n"
	return hi
}

func buildDetailView(conn api.Connection) string {
	output := ""
	output += addDepartureStation(conn.Departure)
	output = addVerticalBar(output, 2)

	for _, stop := range conn.Vias.Via {
		output = addVia(output, stop)
	}

	output += addArrivalStation(conn.Arrival)
	output += "\n"
	return output
}
