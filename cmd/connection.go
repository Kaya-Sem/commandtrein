package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Kaya-Sem/commandtrein/api"
	"github.com/Kaya-Sem/commandtrein/internal/util"
	table "github.com/Kaya-Sem/commandtrein/tables"
	teaTable "github.com/charmbracelet/bubbles/table"
)

const (
	maandag    = "maandag"
	dinsdag    = "dinsdag"
	woensdag   = "woensdag"
	donderdag  = "donderdag"
	vrijdag    = "vrijdag"
	zaterdag   = "zaterdag"
	zondag     = "zondag"
	morgen     = "morgen"
	overmorgen = "overmorgen"
	empty      = ""
)

func getDate(date string) (string, error) {
	if date == empty {
		// current date in ddmmyy
		now := time.Now()
		return fmt.Sprintf("%02d%02d%02d", now.Day(), now.Month(), now.Year()%100), nil
	}

	switch strings.ToLower(date) {
	case morgen:
		tomorrow := time.Now().AddDate(0, 0, 1)
		return fmt.Sprintf("%02d%02d%02d", tomorrow.Day(), tomorrow.Month(), tomorrow.Year()%100), nil
	case overmorgen:
		dayAfterTomorrow := time.Now().AddDate(0, 0, 2)
		return fmt.Sprintf("%02d%02d%02d", dayAfterTomorrow.Day(), dayAfterTomorrow.Month(), dayAfterTomorrow.Year()%100), nil
	case maandag, dinsdag, woensdag, donderdag, vrijdag, zaterdag, zondag:
		return getNextValidDay(date), nil
	}

	// Handle ddmmyy format
	if len(date) == 6 {
		_, err := strconv.Atoi(date)
		if err == nil {
			return date, nil
		}
	}

	return "", fmt.Errorf("invalid date format: %s", date)
}

func getNextValidDay(dayName string) string {
	dayMap := map[string]time.Weekday{
		maandag:   time.Monday,
		dinsdag:   time.Tuesday,
		woensdag:  time.Wednesday,
		donderdag: time.Thursday,
		vrijdag:   time.Friday,
		zaterdag:  time.Saturday,
		zondag:    time.Sunday,
	}

	targetDay, exists := dayMap[strings.ToLower(dayName)]
	if !exists {
		return ""
	}

	now := time.Now()
	currentDay := now.Weekday()

	// Calculate days to add
	daysToAdd := int(targetDay - currentDay)
	if daysToAdd <= 0 {
		// If target day is today or in the past, get next week's occurrence
		daysToAdd += 7
	}

	nextValidDay := now.AddDate(0, 0, daysToAdd)
	return fmt.Sprintf("%02d%02d%02d", nextValidDay.Day(), nextValidDay.Month(), nextValidDay.Year()%100)
}

func handleConnection(stationFrom string, stationTo string, timeQuery string, departure bool, date string) {
	s := util.NewSpinner("", " fetching connections", 1*time.Second)
	s.Start()

	// Parse the date if provided
	parsedDate := ""
	if date != "" {
		var err error
		parsedDate, err = getDate(date)
		if err != nil {
			fmt.Printf("Error parsing date: %v\n", err)
			os.Exit(1)
		}
	}

	connectionsJSON, err := api.GetConnections(stationFrom, stationTo, timeQuery, departure, parsedDate)
	if err != nil {
		panic(err)
	}

	connections, err := api.ParseConnections(connectionsJSON)
	if err != nil {
		panic(err)
	}

	columns := []teaTable.Column{
		{Title: "Vertrek", Width: 10},
		{Title: "Reistijd", Width: 9},
		{Title: "Aankomst", Width: 8},
		{Title: "Spoor", Width: 10},
	}

	rows := make([]teaTable.Row, len(connections))

	for i, conn := range connections {
		departureTimeWithDelay := util.UnixToHHMM(conn.Departure.Time)
		delay := util.FormatDelay(conn.Departure.Delay)
		if delay != "" {
			departureTimeWithDelay += " " + delay
		}

		rows[i] = teaTable.Row{
			departureTimeWithDelay,
			util.GetDurationInMinutes(conn.Duration),
			util.UnixToHHMM(conn.Arrival.Time),
			conn.Departure.Platform,
		}
	}

	s.Stop()
	table.RenderTable(columns, rows, connections)
}
