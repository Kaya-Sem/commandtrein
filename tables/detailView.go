package table

import (
	"fmt"
	"strings"

	"github.com/Kaya-Sem/commandtrein/internal/util"

	"github.com/Kaya-Sem/commandtrein/api"
)

const (
	YELLOW            = "\033[33m"
	RED               = "\033[31m"
	ITALIC            = "\033[3m"
	DIM               = "\033[2m"
	RESET             = "\033[0m"
	verticalBar       = "┃"
	bottomCorner      = "┗━"
	topCorner         = "┏━"
	trackSwitchDotted = "┊"
)

func dim(text string) string {
	return fmt.Sprintf("%s%s%s", DIM, text, RESET)
}

func yellow(text string) string {
	return fmt.Sprintf("%s%s%s", YELLOW, text, RESET)
}

func red(text string) string {
	return fmt.Sprintf("%s%s%s", RED, text, RESET)
}

func italic(text string) string {
	return fmt.Sprintf("%s%s%s", ITALIC, text, RESET)
}

func addVerticalBar(s string, repetitions int) string {
	return s + strings.Repeat(fmt.Sprintf("        %s\n", yellow(verticalBar)), repetitions)
}

func addVia(text string, v api.ViaInfo) string {
	text += fmt.Sprintf(" %s  %s %s, platform %s\n",
		util.UnixToHHMM(v.Arrival.Time),
		yellow(bottomCorner),
		v.Arrival.Station,
		v.Arrival.Platform)
	text += fmt.Sprintf("        %s\n", trackSwitchDotted)
	text += fmt.Sprintf(" %s  %s %s, platform %s\n",
		util.UnixToHHMM(v.Departure.Time),
		yellow(topCorner),
		v.Departure.Station,
		v.Departure.Platform)
	return addVerticalBar(text, 3)
}

func addArrivalStation(a api.ConnectionArrival) string {
	return fmt.Sprintf(" %s  %s %s\n",
		util.UnixToHHMM(a.Time),
		yellow(bottomCorner),
		a.Station)
}

func addDepartureStation(c api.Connection) string {
	delay := util.FormatDelay(c.Departure.Delay)
	paddedDelay := red(RightPad(delay, 3)) // Padding delay to a total width of 7

	// first line: time, yellow top corner, departure station name
	header := fmt.Sprintf(" %s  %s %s ", util.UnixToHHMM(c.Departure.Time), yellow(topCorner), c.Departure.Station)
	header += "\n"

	// second line: delay, vertical bar and relative departure
	header += fmt.Sprintf("   %s  %s  %s", paddedDelay, yellow(verticalBar), dim(italic("vertrekt "+CalculateHumanRelativeTime(c))))

	header += "\n"

	// third line
	header += fmt.Sprintf("        %s", yellow(verticalBar))

	header += "\n"

	return header

}

func buildDetailView(conn api.Connection) string {
	output := addDepartureStation(conn)
	output = addVerticalBar(output, 4)

	for _, stop := range conn.Vias.Via {
		output = addVia(output, stop)
	}
	output += addArrivalStation(conn.Arrival)
	return output + "\n"
}
