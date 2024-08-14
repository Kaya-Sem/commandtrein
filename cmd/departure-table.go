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
	BorderForeground(lipgloss.Color("240"))

var delayStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("9")). // Red color
	Italic(true)

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[0]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func printDepartureTable(departures []timetableDeparture) {
	fmt.Println()
	columns := []table.Column{
		{Title: "Destination", Width: 20},
		{Title: "Time", Width: 10},
		{Title: "Delay", Width: 5},
		{Title: "Track", Width: 10},
	}

	var rows []table.Row
	for _, dep := range departures {

		time, err := strconv.ParseInt(dep.Time, 10, 64)
		if err != nil {
			fmt.Printf(`Could not convert string %s`, dep.Time)
		}
		departureTime := UnixToHHMM(time)
		if err != nil {
			fmt.Printf("Error converting time: %v\n", err)
			return
		}

		// delay is represented as seconds.
		delay, _ := strconv.Atoi(dep.Delay)
		delay = delay / 60
		formattedDelay := strconv.Itoa(delay)
		if delay > 0 {
			formattedDelay = "+" + FormatDelay(delay)
		}
		row := table.Row{
			dep.Station,
			departureTime,
			formattedDelay,
			dep.Platform,
		}

		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(19),
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

	m := model{t}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
