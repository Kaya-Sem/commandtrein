package table

import (
	"fmt"
	"os"

	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TableData interface {
	api.Connection | api.TimetableDeparture
}

type TableModel[T TableData] struct {
	table           table.Model
	selectedDetails string
	showMessage     bool
	message         string
	data            []T
}

func (m *TableModel[T]) Init() tea.Cmd { return nil }

func (m *TableModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.showMessage = true
				m.message = m.getDetailedInfo(m.data[selectedIndex])
			}
			return m, tea.Quit
		}
	}
	m.table, teaCmd = m.table.Update(msg)
	m.updateSelectedDetails()
	return m, teaCmd
}

func (m *TableModel[T]) View() string {
	if m.showMessage {
		return m.message
	}
	tableView := m.table.View()
	detailsView := detailsBoxStyle.Render(m.selectedDetails)
	return lipgloss.JoinHorizontal(lipgloss.Top, tableView, detailsView)
}

func (m *TableModel[T]) updateSelectedDetails() {
	selectedRow := m.table.SelectedRow()
	if selectedRow != nil {
		selectedIndex := m.table.Cursor()
		m.selectedDetails = m.getDetailedInfo(m.data[selectedIndex])
	} else {
		m.selectedDetails = "Nothing found!"
	}
}

func (m *TableModel[T]) getDetailedInfo(item T) string {
	switch v := any(item).(type) {
	case api.Connection:
		return getDetailedConnectionInfo(v)
	case api.TimetableDeparture:
		return getDetailedDepartureInfo(v)
	default:
		return "Unsupported data type"
	}
}

func RenderTable[T TableData](
	columnItems []table.Column,
	rowItems []table.Row,
	data []T,
) {
	fmt.Println()
	t := table.New(
		table.WithColumns(columnItems),
		table.WithRows(rowItems),
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
	m := &TableModel[T]{
		table: t,
		data:  data,
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
