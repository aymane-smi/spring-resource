package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string
type itemDelegate struct{}

type model struct {
	choices   list.Model
	choice    any `default:"0"`
	quitting  bool
	typeModel int
}

var (
	//titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(0)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("170"))
	//paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	//helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func (i item) FilterValue() string { return "" }

func (m model) View() string {
	if m.choice != nil {
		if m.typeModel == 2 {
			return quitTextStyle.Render(fmt.Sprintf("you choiced %s", m.choice))
		} else if m.typeModel == 1 {
			return quitTextStyle.Render(fmt.Sprintf("you choiced %d", m.choice))
		}
	}
	if m.quitting {
		return quitTextStyle.Render("exit?, no more generation now.")
	}
	return "\n" + m.choices.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.choices.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.choices.SelectedItem().(item)
			if ok && m.typeModel == 1 {
				switch string(i) {
				case "Integer":
					m.choice = 1
				case "UUID":
					m.choice = 3
				case "Long":
					m.choice = 2
				}
			} else if ok && m.typeModel == 2 {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.choices, cmd = m.choices.Update(msg)
	return m, cmd
}
func (m model) Init() tea.Cmd {
	return nil
}

// option 1 for entity generation
// option 2 for repository generation
func Execute(option int) (any, error) {
	var items []list.Item
	if option == 1 {
		items = []list.Item{
			item("Integer"),
			item("UUID"),
			item("Long"),
		}
	} else if option == 2 {
		items = []list.Item{
			item("JpaRepository"),
			item("CrudRepository"),
			item("Repository"),
		}
	}
	l := list.New(items, itemDelegate{}, 20, 14)
	l.Title = "test"
	if option == 1 {
		l.Title = "Choose entity id type?"
	} else if option == 2 {
		l.Title = "Choose repository interface type?"
	}
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	m := model{choices: l, typeModel: option}

	if result, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	} else {
		return result.(model).choice, nil
	}

	return 0, errors.New("error during list generation")
}
