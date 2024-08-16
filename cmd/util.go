package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func normalizeTime(time string) string {
	time = strings.ReplaceAll(time, " ", "")
	time = strings.ReplaceAll(time, ":", "")
	return time
}

func UnixToHHMM(unixTime string) string {
	unixTimeInt, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		fmt.Printf("could not parse timestring: %s", unixTime)
		return "99:99"
	}
	t := time.Unix(unixTimeInt, 0).Local()
	return t.Format("15:04")
}

func FormatDelay(minutes int) string {
	if minutes >= 60 {
		hours := minutes / 60
		remainingMinutes := minutes % 60
		if remainingMinutes > 0 {
			return strconv.Itoa(hours) + "h " + strconv.Itoa(remainingMinutes) + "m"
		}
		return strconv.Itoa(hours) + "h"
	}
	return strconv.Itoa(minutes)
}

func ShiftArgs(args []string) []string {
	return args[1:]
}

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
	} else {
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		if minutes == 0 {
			return fmt.Sprintf("%d hours", hours)
		}
		return fmt.Sprintf("%d hours %d min", hours, minutes)
	}
}
