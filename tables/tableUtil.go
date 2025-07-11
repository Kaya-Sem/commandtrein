package table

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
)

const (
	Gray               = "240"
	White              = "15"
	Green              = "2"
	Orange             = "214"
	Red                = "9"
	BorderColor        = Gray
	SelectedForeground = White
	SelectedBackground = "#006ab3" // SNCB blue
	tableHeight        = 15
)

var DetailsBoxStyle = lipgloss.NewStyle().Padding(1)

type timeable interface {
	GetUnixDepartureTime() int
	GetDelayInSeconds() int
}

// CalculateHumanRelativeTime returns a human-readable time difference between
// the current time and a given timeable event, such as a train departure.
//
// Parameters:
//   - t (timeable): an object with methods GetUnixDepartureTime() and GetDelayInSeconds().
// Returns:
//   - string: the formatted time difference.

func CalculateHumanRelativeTime(t timeable) string {
	now := time.Now()

	depTime := time.Unix(int64(t.GetUnixDepartureTime()), 0)
	depTime = depTime.Add(time.Duration(t.GetDelayInSeconds()) * time.Second)

	// Calculate the duration between now and the adjusted departure time
	duration := depTime.Sub(now)

	if duration < 1*time.Minute {
		return "zometeen"
	}

	if duration < 60*time.Minute {
		return fmt.Sprintf("over %d minuten", int(duration.Minutes()))
	}

	// less than 2 hours
	if duration < 120*time.Minute {
		minutes := int(duration.Minutes()) % 60
		if minutes == 0 {
			return "over 1u"
		}

		var minuteString = fmt.Sprintf("%d", minutes)
		if minutes < 10 {
			minuteString = "0" + minuteString
		}

		return fmt.Sprintf("over 1u%sm", minuteString)
	}

	// Check if it's over 24 hours
	if duration >= 24*time.Hour {
		days := int(duration.Hours()) / 24
		if days == 1 {
			return "over 1 dag"
		}
		return fmt.Sprintf("over %d dagen", days)
	}

	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	if minutes == 0 {
		return fmt.Sprintf("over %du", hours)
	}

	return fmt.Sprintf("over %du%d", hours, minutes)
}

func LeftPad(s string, padWidth int) string {
	padding := padWidth - len(s)
	if padding > 0 {
		return strings.Repeat(" ", padding) + s
	}

	return s
}

func RightPad(s string, padWidth int) string {
	padding := padWidth - len(s)
	if padding > 0 {
		return s + strings.Repeat(" ", padding)
	}

	return s
}
