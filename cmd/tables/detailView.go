package table

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"strings"

	"github.com/Kaya-Sem/commandtrein/cmd/api"
)

const (
	yellowCode        = "\033[33m"
	redCode           = "\033[31m"
	italicCode        = "\033[3m"
	dimCode           = "\033[2m"
	resetCode         = "\033[0m"
	verticalBar       = "┃"
	bottomCorner      = "┗━"
	topCorner         = "┏━"
	trackSwitchDotted = "┊"
)

func dim(text string) string {
	return fmt.Sprintf("%s%s%s", dimCode, text, resetCode)
}

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

func addDepartureStation(c api.Connection) string {
	delay := cmd.FormatDelay(c.Departure.Delay)
	paddedDelay := RightPad(red(delay), 11) // Padding delay to a total width of 7

	return fmt.Sprintf(" %s  %s %s \n    %s  %s  %s\n",
		cmd.UnixToHHMM(c.Departure.Time),
		yellow(topCorner),
		c.Departure.Station,
		paddedDelay,
		yellow(verticalBar),
		dim(italic("departure in "+CalculateHumanRelativeTime(c))),
	)
}

func buildDetailView(conn api.Connection) string {
	output := addDepartureStation(conn)
	output = addVerticalBar(output, 3)

	for _, stop := range conn.Vias.Via {
		output = addVia(output, stop)
	}
	output += addArrivalStation(conn.Arrival)
	return output + "\n"
}
