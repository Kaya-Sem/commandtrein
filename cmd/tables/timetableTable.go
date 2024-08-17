package table

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	lowOccupancyStyle    = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("2"))   // green
	mediumOccupancyStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("214")) // orange
	highOccupancyStyle   = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("9"))   // red

)

func styleOccupancy(s string) string {
	if s == "low" {
		return lowOccupancyStyle.Render(s)
	}
	if s == "medium" {
		return mediumOccupancyStyle.Render(s)
	}

	return highOccupancyStyle.Render(s)
}

type timetableTableModel struct {
	table           table.Model
	selectedDetails string
	departures      []api.TimetableDeparture
}

func (m timetableTableModel) Init() tea.Cmd { return nil }

func getDetailedDepartureInfo(d api.TimetableDeparture, relativeTime string) string {
	return fmt.Sprintf(`
Departure in: %s
Track: %s
Departure Time: %s
Vehicle: %s
Occupancy: %s
`,
		relativeTime,
		d.Platform,
		cmd.UnixToHHMM(d.Time),
		d.Vehicle,
		styleOccupancy(d.Occupancy.Name),
	)
}

func (m *timetableTableModel) updateSelectedDetails() {
	selectedRow := m.table.SelectedRow()
	if selectedRow != nil {
		selectedIndex := m.table.Cursor()
		selectedDeparture := m.departures[selectedIndex]

		// Calculate the relative time for the selected row
		departureTime := selectedRow[0]
		relativeTime := CalculateHumanRelativeTime(departureTime)

		// Update the selected details including the relative time
		m.selectedDetails = getDetailedDepartureInfo(selectedDeparture, relativeTime)
	} else {
		m.selectedDetails = "No row selected" // Should never really happen
	}
}

func (m timetableTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var teaCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	m.table, teaCmd = m.table.Update(msg)

	m.updateSelectedDetails()

	return m, teaCmd
}

var detailsBoxStyle = lipgloss.NewStyle().Padding(1) //.Border(lipgloss.NormalBorder())

func (m timetableTableModel) View() string {
	tableView := m.table.View()
	detailsView := detailsBoxStyle.Render(m.selectedDetails)

	return lipgloss.JoinHorizontal(lipgloss.Top, tableView, detailsView)
}

func RenderTimetableTable(
	columnItems []table.Column,
	rowItems []table.Row,
	departures []api.TimetableDeparture,
) {
	fmt.Println()

	columns := columnItems
	rows := rowItems

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(tableHeight),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(BorderColor)).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color(SelectedForeground)).
		Background(lipgloss.Color(SelectedBackground))
	t.SetStyles(s)

	m := timetableTableModel{
		table:      t,
		departures: departures,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
