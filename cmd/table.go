package cmd

import (
	"fmt"
	"os"
	"strconv"

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
		relativeTime := CalculateHumanRelativeTime(departureTime)
		m.relativeTime = relativeTime
	} else {
		m.relativeTime = ""
	}

	return m, cmd
}

func (m model) View() string {
	// Add the relative time to the view only if there is a selected row
	if m.relativeTime != "" {
		return baseStyle1.Render(m.table.View()) + "\n\n" + "Departure in: " + m.relativeTime + "\n"
	}
	return baseStyle1.Render(m.table.View()) + "\n"
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
		departureTime := UnixToHHMM(connection.Departure.Time)
		arrivalTime := UnixToHHMM(connection.Arrival.Time)

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
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t, ""}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
