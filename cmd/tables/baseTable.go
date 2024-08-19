package table

import (
	"fmt"
	"os"

	"github.com/Kaya-Sem/commandtrein/cmd/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Data interface {
	api.Connection | api.TimetableDeparture
}

type Model[T Data] struct {
	table           table.Model
	selectedDetails string
	showMessage     bool
	message         string
	data            []T
}

func (m *Model[T]) Init() tea.Cmd { return nil }

func (m *Model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *Model[T]) View() string {
	if m.showMessage {
		return m.message
	}

	// Width terminal
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))

	dynamicHeight := len(m.data)
	if dynamicHeight > 5 {
		dynamicHeight = 5
	}

	if lipgloss.Width(m.table.View())+lipgloss.Width(m.selectedDetails) <= width {
		// Horizontal layout
		m.table.SetHeight(tableHeight)
	} else {
		// Vertical layout
		m.table.SetHeight(dynamicHeight)
	}

	tableView := m.table.View()
	detailsView := DetailsBoxStyle.Render(m.selectedDetails)

	if lipgloss.Width(tableView)+lipgloss.Width(detailsView) <= width {
		return lipgloss.JoinHorizontal(lipgloss.Top, tableView, detailsView)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left, tableView, detailsView)
	}
}

func (m *Model[T]) updateSelectedDetails() {
	selectedRow := m.table.SelectedRow()
	if selectedRow != nil {
		selectedIndex := m.table.Cursor()
		m.selectedDetails = m.getDetailedInfo(m.data[selectedIndex])
	} else {
		m.selectedDetails = "Nothing found!"
	}
}

// TODO: do this with an interface instead
func (m *Model[T]) getDetailedInfo(item T) string {
	switch v := any(item).(type) {
	case api.Connection:
		return getDetailedConnectionInfo(v)
	case api.TimetableDeparture:
		return getDetailedDepartureInfo(v)
	default:
		return "Unsupported data type"
	}
}

func RenderTable[T Data](
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
	m := &Model[T]{
		table: t,
		data:  data,
	}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
