package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(0, 0)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type responseMsg struct{}

type item struct {
	title       string
	description string
	appName     string
	id          int
}

func (i item) Title() string {
	if i.appName != "" {
		return i.title + " (" + i.appName + ")"
	}
	return i.title
}
func (i item) Description() string { return i.description }
func (i item) Id() int             { return i.id }
func (i item) FilterValue() string { return i.title + i.description }

type listKeyMap struct {
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type model struct {
	sub          chan struct{}
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func getItems() []list.Item {
	dunstctlHistory := getDunstctlHistory()
	var items []list.Item
	for _, it := range dunstctlHistory {
		items = append(items, item{
			title:       it.Summary.Data,
			description: it.Body.Data,
			id:          it.ID.Data,
			appName:     it.Appname.Data,
		})
	}
	return items
}

func newModel() model {
	var (
		delegateKeys = newDelegateKeyMap()
		listKeys     = newListKeyMap()
	)

	// Make initial list of items
	items := getItems()

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	notificationList := list.New(items, delegate, 0, 0)
	notificationList.Title = "Notifications"
	notificationList.Styles.Title = titleStyle
	notificationList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}

	return model{
		sub:          make(chan struct{}),
		list:         notificationList,
		keys:         listKeys,
		delegateKeys: delegateKeys,
	}
}

func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		cmd := exec.Command("dbus-monitor", "path='/org/freedesktop/Notifications',type='signal',interface='org.freedesktop.DBus.Properties',member='PropertiesChanged'")
		stdout, _ := cmd.StdoutPipe()
		scanner := bufio.NewScanner(stdout)
		go func() {
			for scanner.Scan() {
				// TODO: throttle
				sub <- struct{}{}
			}
		}()
		cmd.Start()
		return cmd.Wait()
	}
}

func waitForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForActivity(m.sub),
		waitForActivity(m.sub),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case responseMsg:
		newItems := getItems()
		m.list.SetItems(newItems)
		return m, waitForActivity(m.sub)

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if _, err := tea.NewProgram(newModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
