package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TODO: needed for flag parsing
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

func FormatDelay(seconds string) string {
	minutes, err := strconv.Atoi(seconds)
	if err != nil {
		return "err"
	}

	minutes /= 60

	// If the delay is 60 minutes or more, convert to hours and minutes
	if minutes >= 60 {
		hours := minutes / 60
		remainingMinutes := minutes % 60
		if remainingMinutes > 0 {
			return "+" + strconv.Itoa(hours) + "h " + strconv.Itoa(remainingMinutes) + "m"
		}
		return "+" + strconv.Itoa(hours) + "h"
	}

	return "+" + strconv.Itoa(minutes) + "m"
}

func ShiftArgs(args []string) []string {
	return args[1:]
}
