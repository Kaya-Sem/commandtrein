package table

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	Gray               = "240"
	White              = "15"
	Green              = "2"
	Orange             = "214"
	Mauve              = "97"
	Red                = "9"
	BorderColor        = Gray
	SelectedForeground = White
	SelectedBackground = Mauve
	tableHeight        = 10
)

var DetailsBoxStyle = lipgloss.NewStyle().Padding(1)

var (
	OccupancyStyle = lipgloss.NewStyle().Italic(true)

	lowOccupancyStyle     = OccupancyStyle.Copy().Foreground(lipgloss.Color(Green))
	mediumOccupancyStyle  = OccupancyStyle.Copy().Foreground(lipgloss.Color(Orange))
	highOccupancyStyle    = OccupancyStyle.Copy().Foreground(lipgloss.Color(Red))
	unknownOccupancyStyle = OccupancyStyle.Copy().Faint(true)
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
