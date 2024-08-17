package table

import (
	"fmt"
	"time"
)

const (
	BorderColor        = "240" // gray
	SelectedForeground = "229" // not setting it to yellow will make the text yellow (yellow on purple = white?)
	SelectedBackground = "57"  // purple
	tableHeight        = 10
)

// CalculateHumanRelativeTime used for calucating human-readable "from now" time. E.g 'in 20 minutes'
func CalculateHumanRelativeTime(departureTime string) string {
	now := time.Now()

	depTime, err := time.Parse("15:04", departureTime)
	if err != nil {
		return ""
	}

	// Combine the parsed time with today's date
	depDateTime := time.Date(now.Year(), now.Month(), now.Day(), depTime.Hour(), depTime.Minute(), 0, 0, now.Location())

	// If the departure time is earlier than now, assume it's for the next day
	if depDateTime.Before(now) {
		depDateTime = depDateTime.Add(24 * time.Hour)
	}

	// Calculate the duration between now and the departure time
	duration := depDateTime.Sub(now)

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
