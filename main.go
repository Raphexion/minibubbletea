package main

// A minimal bubbletea program. Quit by pressing "q"

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	X bool
}

type tickMsg struct{}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.X = m.X != true
		return m, tick()
	}

	return m, nil
}

func (m model) View() string {
	if m.X {
		return "X"
	}
	return " "
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func main() {
	p := tea.NewProgram(model{true})
	p.Start()
}
