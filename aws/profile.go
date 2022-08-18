package sso

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nanih98/aws-sso/logger"
	"github.com/nanih98/aws-sso/utils"
)

const listHeight = 30

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type model struct {
	list     list.Model
	items    []item
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s used as AWS_PROFILE :)", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("AWS PROFILE NOT SELECTED.")
	}
	return "\n" + m.list.View()
}

func GetProfiles(filter, filepath string) []string {
	b, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
	}

	data := regexp.MustCompile(`\w+-` + filter)
	//data := regexp.MustCompile(`\[([^\[\]]*)\]`) // prints content inside brackets, without filtering

	profiles := data.FindAllString(string(b), -1)

	return profiles
}

func Profile(log *logger.CustomLogger, filter string) {
	log.Info("Reading file .aws/credentials")
	log.Info(fmt.Sprintf("Serching credentials for profile name %s", filter))

	credentialsPath := utils.GetUserHome(log) + "/.aws/credentials"

	if filter == "" {
		log.Info("Searching without filter")
	}

	profiles := GetProfiles(filter, credentialsPath) //this is a harcoded path

	var items []list.Item

	for _, profile := range profiles {
		profile = strings.Trim(profile, "[")
		profile = strings.Trim(profile, "]")
		items = append(items, item(profile))
	}

	if len(items) == 0 {
		log.Fatal(fmt.Errorf("Profiles not found for %s", filter))
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Select your AWS PROFILE"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
