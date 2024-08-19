package table

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	Gray               = "240"
	White              = "15"
	Mauve              = "97"
	BorderColor        = Gray
	SelectedForeground = White
	SelectedBackground = Mauve
	tableHeight        = 10
)

var DetailsBoxStyle = lipgloss.NewStyle().Padding(1)

var (
	baseOccupancyStyle = lipgloss.NewStyle().Italic(true)

	lowOccupancyStyle     = baseOccupancyStyle.Copy().Foreground(lipgloss.Color("2"))   // green
	mediumOccupancyStyle  = baseOccupancyStyle.Copy().Foreground(lipgloss.Color("214")) // orange
	highOccupancyStyle    = baseOccupancyStyle.Copy().Foreground(lipgloss.Color("9"))   // red
	unknownOccupancyStyle = baseOccupancyStyle.Copy().Faint(true)
)

func styleOccupancy(s string) string {
	switch s {
	case "low":
		return lowOccupancyStyle.Render(s)
	case "medium":
		return mediumOccupancyStyle.Render(s)
	case "high":
		return highOccupancyStyle.Render(s)
	default:
		return unknownOccupancyStyle.Render(s)
	}
}

type timeable interface {
	GetUnixDepartureTime() int
	GetDelayInSeconds() int
}

func CalculateHumanRelativeTime(t timeable) string {
	now := time.Now()

	depTime := time.Unix(int64(t.GetUnixDepartureTime()), 0)
	depTime = depTime.Add(time.Duration(t.GetDelayInSeconds()) * time.Second)

	// Calculate the duration between now and the adjusted departure time
	duration := depTime.Sub(now)

	// Handle special cases
	if duration < 1*time.Minute {
		return "now"
	} else if duration < 60*time.Minute {
		return fmt.Sprintf("%d min", int(duration.Minutes()))
	} else if duration < 120*time.Minute {
		minutes := int(duration.Minutes()) % 60
		if minutes == 0 {
			return "1 hour"
		}
		return fmt.Sprintf("1 hour %d min", minutes)
	}

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	if minutes == 0 {
		return fmt.Sprintf("%d hours", hours)
	}

	return fmt.Sprintf("%d hours %d min", hours, minutes)
}
