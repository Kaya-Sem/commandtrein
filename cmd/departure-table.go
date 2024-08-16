package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle1 = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("9"))

type tableModel struct {
	table        table.Model
	relativeTime string
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
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

func (m tableModel) View() string {
	// Add the relative time to the view only if there is a selected row
	if m.relativeTime != "" {
		return baseStyle1.Render(m.table.View()) + "\n\n" + "Departure in: " + m.relativeTime + "\n"
	}
	return baseStyle1.Render(m.table.View()) + "\n"
}

func RenderTable(columnItems []table.Column, rowItems []table.Row) {
	fmt.Println()

	columns := columnItems
	rows := rowItems

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(6),
	)

	s := table.DefaultStyles()
	//	s.Cell.Align(lipgloss.Position(4))
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57"))
	t.SetStyles(s)

	m := tableModel{t, ""}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
