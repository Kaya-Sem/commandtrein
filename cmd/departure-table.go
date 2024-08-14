package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("9"))

type model struct {
	table        table.Model
	relativeTime string
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
			// case "enter":
			// 	return m, tea.Batch(
			// 		tea.Printf("Let's go to %s!", m.table.SelectedRow()[0]),
			// 	)
		}
	}

	m.table, cmd = m.table.Update(msg)

	// Calculate the relative time for the currently selected row
	selectedRow := m.table.SelectedRow()
	if selectedRow != nil {
		departureTime := selectedRow[0]
		relativeTime := calculateRelativeTime(departureTime)
		m.relativeTime = relativeTime
	} else {
		m.relativeTime = ""
	}

	return m, cmd
}

func (m model) View() string {
	// Add the relative time to the view only if there is a selected row
	if m.relativeTime != "" {
		return baseStyle.Render(m.table.View()) + "\n\n" + "Departure in: " + m.relativeTime + "\n"
	}
	return baseStyle.Render(m.table.View()) + "\n"
}

func PrintDepartureTable(connections []Connection) {
	fmt.Println()
	columns := []table.Column{
		{Title: "Departure", Width: 10},
		{Title: "Duration", Width: 10},
		{Title: "Arrival", Width: 10},
		{Title: "Track", Width: 5},
	}

	var rows []table.Row
	for _, connection := range connections {

		departureTimeInt, err := strconv.ParseInt(connection.Departure.Time, 10, 64)
		if err != nil {
			fmt.Printf(`Could not convert string %s`, connection.Departure.Time)
		}

		arrivalTimeInt, err := strconv.ParseInt(connection.Arrival.Time, 10, 64)
		if err != nil {
			fmt.Printf(`Could not convert string %s`, connection.Departure.Time)
		}
		departureTime := UnixToHHMM(departureTimeInt)
		if err != nil {
			fmt.Printf("Error converting time: %v\n", err)
			return
		}

		arrivalTime := UnixToHHMM(arrivalTimeInt)
		if err != nil {
			fmt.Printf("Error converting time: %v\n", err)
			return
		}

		// delay is represented as seconds.
		delay, _ := strconv.Atoi(connection.Departure.Delay)
		delay = delay / 60
		// formattedDelay := strconv.Itoa(delay)
		// if delay > 0 {
		// 	formattedDelay = "+" + FormatDelay(delay)
		// }

		durationInt, _ := strconv.ParseInt(connection.Duration, 10, 32)

		duration := strconv.FormatInt(durationInt/60, 10) + "m"

		row := table.Row{
			departureTime,
			duration,
			arrivalTime,
			connection.Departure.Platform,
		}

		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(6),
	)

	s := table.DefaultStyles()
	s.Cell.Align(lipgloss.Position(4))
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	// s.Selected = s.Selected.
	// 	Foreground(lipgloss.Color("229")).
	// 	Background(lipgloss.Color("57")).
	// 	Bold(false)
	t.SetStyles(s)

	m := model{t, ""}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func calculateRelativeTime(departureTime string) string {
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
