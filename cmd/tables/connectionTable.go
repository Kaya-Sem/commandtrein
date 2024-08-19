package table

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
)

func getDetailedConnectionInfo(c api.Connection) string {
	return fmt.Sprintf(`
	Departure in %s
	Destination: %s
	Track: %s
	Departure Time: %s
	Vehicle: %s
`,
		CalculateHumanRelativeTime(c),
		c.Departure.Station,
		c.Departure.Platform,
		cmd.UnixToHHMM(c.Departure.Time),
		c.Departure.Vehicle,
	)
}
