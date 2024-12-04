package util

import (
	"fmt"
	"strconv"
	"time"
)

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

	if minutes == 0 {
		return ""
	}

	if minutes >= 60 {
		hours := minutes / 60
		remainingMinutes := minutes % 60
		if remainingMinutes > 0 {
			return "+" + strconv.Itoa(hours) + "h " + strconv.Itoa(remainingMinutes) + "m"
		}
		return "+" + strconv.Itoa(hours) + "h"
	}

	return "+" + strconv.Itoa(minutes)
}
