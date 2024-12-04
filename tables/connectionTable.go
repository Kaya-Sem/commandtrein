package table

import (
	"fmt"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/Kaya-Sem/commandtrein/internal/util"
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
		util.UnixToHHMM(c.Departure.Time),
		c.Departure.Vehicle,
	)
}
