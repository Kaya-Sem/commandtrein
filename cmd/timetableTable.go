package cmd

import (
	"fmt"
	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	BorderColor        = "240" // gray
	SelectedForeground = "229" // not setting it to yellow will make the text yellow
	SelectedBackground = "57"  // purple
	tableHeight        = 6
)

var baseStyle1 = lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("9"))

type timetableTableModel struct {
	table        table.Model
	relativeTime string
	showMessage  bool
	message      string
	departures   []api.TimetableDeparture
}

func (m timetableTableModel) Init() tea.Cmd { return nil }

func getDetailedDepartureInfo(d api.TimetableDeparture) string {
	return fmt.Sprintf(`
Detailed info:
Destination: %s
Track: %s
Departure Time: %s
Vehicle: %s
Occupancy: %s
`,
		d.Station,
		d.Platform,
		UnixToHHMM(d.Time),
		d.Vehicle,
		d.Occupancy,
	)
}

func (m timetableTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			selectedRow := m.table.SelectedRow()
			if selectedRow != nil {
				selectedIndex := m.table.Cursor()
				selectedDeparture := m.departures[selectedIndex]
				m.showMessage = true
				m.message = getDetailedDepartureInfo(selectedDeparture)
			}
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

var italicStyle = lipgloss.NewStyle().Italic(true)

func (m timetableTableModel) View() string {
	if m.showMessage {
		// Show the message instead of the table if the flag is set
		return baseStyle1.Render(m.message)
	}

	// Add the relative time to the view only if there is a selected row
	if m.relativeTime != "" {
		return baseStyle1.Render(m.table.View()) + "\n\n" + "Departure in: " + italicStyle.Render(m.relativeTime) + "\n"
	}
	return baseStyle1.Render(m.table.View()) + "\n"
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
		departures: departures, // Store the departures
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
