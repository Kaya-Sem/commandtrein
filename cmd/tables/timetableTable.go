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

type timetableTableModel struct {
	table           table.Model
	selectedDetails string
	showMessage     bool
	message         string
	departures      []api.TimetableDeparture
}

func (m timetableTableModel) Init() tea.Cmd { return nil }

func getDetailedDepartureInfo(d api.TimetableDeparture) string {
	relativeTime := CalculateHumanRelativeTime(d)
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

		// Update the selected details including the relative time
		m.selectedDetails = getDetailedDepartureInfo(selectedDeparture)
	} else {
		m.selectedDetails = "Nothing found!"
	}
}

func (m timetableTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var teaCmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			selectedRow := m.table.SelectedRow()
			if selectedRow != nil {
				selectedIndex := m.table.Cursor
				selectedDeparture := m.departures[selectedIndex()]
				m.showMessage = true
				m.message = getDetailedDepartureInfo(selectedDeparture)
			}
			return m, tea.Quit
		}
	}

	m.table, teaCmd = m.table.Update(msg)

	m.updateSelectedDetails()

	return m, teaCmd
}

// TODO: export to consts
var detailsBoxStyle = lipgloss.NewStyle().Padding(1)

func (m timetableTableModel) View() string {
	if m.showMessage {
		return m.message
	}
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
