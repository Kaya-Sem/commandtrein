package table

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"github.com/charmbracelet/lipgloss"
)

func getDetailedDepartureInfo(d api.TimetableDeparture) string {
	relativeTime := CalculateHumanRelativeTime(d)
	return fmt.Sprintf(`
Departure in: %s
Track: %s
Departure Time: %s
Vehicle: %s
Occupancy: %s
`,
		relativeTime,
		d.Platform,
		cmd.UnixToHHMM(d.Time),
		d.Vehicle,
		styleOccupancy(d.Occupancy.Name),
	)
}

// TODO: export to consts
var detailsBoxStyle = lipgloss.NewStyle().Padding(1)
