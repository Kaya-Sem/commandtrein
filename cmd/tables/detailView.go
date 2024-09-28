package table

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"strings"
)

const (
	yellowCode        = "\033[33m"
	redCode           = "\033[31m"
	italicCode        = "\033[3m"
	resetCode         = "\033[0m"
	verticalBar       = "┃"
	bottomCorner      = "┗━"
	topCorner         = "┏━"
	trackSwitchDotted = "┊"
)

func yellow(text string) string {
	return fmt.Sprintf("%s%s%s", yellowCode, text, resetCode)
}

func red(text string) string {
	return fmt.Sprintf("%s%s%s", redCode, text, resetCode)
}

func italic(text string) string {
	return fmt.Sprintf("%s%s%s", italicCode, text, resetCode)
}

func addVerticalBar(s string, repetitions int) string {
	return s + strings.Repeat(fmt.Sprintf("        %s\n", yellow(verticalBar)), repetitions)
}

func addVia(text string, v api.ViaInfo) string {
	text += fmt.Sprintf(" %s  %s %s, platform %s\n",
		cmd.UnixToHHMM(v.Arrival.Time),
		yellow(bottomCorner),
		v.Arrival.Station,
		v.Arrival.Platform)
	text += fmt.Sprintf("        %s\n", trackSwitchDotted)
	text += fmt.Sprintf(" %s  %s %s, platform %s\n",
		cmd.UnixToHHMM(v.Departure.Time),
		yellow(topCorner),
		v.Departure.Station,
		v.Departure.Platform)
	return addVerticalBar(text, 3)
}

func addArrivalStation(a api.ConnectionArrival) string {
	return fmt.Sprintf(" %s  %s %s\n",
		cmd.UnixToHHMM(a.Time),
		yellow(bottomCorner),
		a.Station)
}

func leftPad(s string, padWidth int) string {
	padding := padWidth - len(s)
	if padding > 0 {
		return s + strings.Repeat(" ", padding)
	}

	return s
}

func addDepartureStation(d api.ConnectionDeparture) string {
	delay := cmd.FormatDelay(d.Delay)
	paddedDelay := leftPad(red(delay), 11) // Padding delay to a total width of 7

	return fmt.Sprintf(" %s  %s %s\n    %s  %s  %s\n",
		cmd.UnixToHHMM(d.Time),
		yellow(topCorner),
		d.Station,
		paddedDelay,
		yellow(verticalBar),
		italic(d.VehicleInfo.ShortName))
}

func buildDetailView(conn api.Connection) string {
	output := addDepartureStation(conn.Departure)
	output = addVerticalBar(output, 2)

	for _, stop := range conn.Vias.Via {
		output = addVia(output, stop)
	}
	output += addArrivalStation(conn.Arrival)
	return output + "\n"
}
