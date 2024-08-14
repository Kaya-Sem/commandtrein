package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	StatusOK                  = 200
	StatusInternalServerError = 500
)

const (
	ErrCli        = 1
	ErrFileRead   = 2
	ErrUnmarshal  = 5
	ErrFileExists = 3
)

var StatusCodes = map[int]string{
	StatusOK:                  "\033[32m200 OK\033[0m",                    // green
	StatusInternalServerError: "\033[31m500 Internal Server Error\033[0m", // red
}

func replaceSpacesWithURLCode(input string) string {
	return strings.ReplaceAll(input, " ", "%20")
}

func getCurrentTimeHHMM() string {
	hours, minutes, _ := time.Now().Clock()
	return fmt.Sprintf("%d%02d", hours, minutes)

}

func normalizeTime(time string) string {
	time = strings.ReplaceAll(time, " ", "")
	time = strings.ReplaceAll(time, ":", "")
	return time
}

func UnixToHHMM(unixTime int64) string {
	t := time.Unix(unixTime, 0).Local()
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
