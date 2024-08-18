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

type connectionTableModel struct {
	table           table.Model
	selectedDetails string
	showMessage     bool
	message         string
	departures      []api.Connection
}

func (m connectionTableModel) Init() tea.Cmd { return nil }

func getDetailedConnectionInfo(c api.Connection) string {
	return fmt.Sprintf(`
Detailed info:
Destination: %s
Track: %s
Departure Time: %s
Vehicle: %s
`,
		c.Departure.Station,
		c.Departure.Platform,
		cmd.UnixToHHMM(c.Departure.Time),
		c.Departure.Vehicle,
	)
}

func (m *connectionTableModel) updateSelectedDetails() {
	selectedRow := m.table.SelectedRow()
	if selectedRow != nil {
		selectedIndex := m.table.Cursor()
		selectedConnection := m.departures[selectedIndex]

		m.selectedDetails = getDetailedConnectionInfo(selectedConnection)
	} else {
		m.selectedDetails = "No row selected" // Should never really happen
	}
}

func (m connectionTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var teaCmd tea.Cmd
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
				m.message = getDetailedConnectionInfo(selectedDeparture)
			}
			return m, tea.Quit
		}
	}

	m.table, teaCmd = m.table.Update(msg)
	m.updateSelectedDetails()

	return m, teaCmd
}

func (m connectionTableModel) View() string {
	if m.showMessage {
		return m.message
	}
	tableView := m.table.View()
	detailsView := detailsBoxStyle.Render(m.selectedDetails)

	return lipgloss.JoinHorizontal(lipgloss.Top, tableView, detailsView)
}

func RenderConnectionTable(
	columnItems []table.Column,
	rowItems []table.Row,
	connections []api.Connection,
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

	m := connectionTableModel{
		table:      t,
		departures: connections,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
