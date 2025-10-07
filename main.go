package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

var globalErr error = nil

type model struct {
	smBanner []string
	mdBanner []string
	lgBanner []string
	width    int
	heigth   int
}

func newSamurai() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return loadAscii
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loadAsciiMsg:
		if msg.err != nil {
			globalErr = msg.err
			return m, tea.Quit
		}
		m.smBanner = msg.sm
		m.mdBanner = msg.md
		m.lgBanner = msg.lg

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.heigth = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return "hi"
}

func main() {
	p := tea.NewProgram(newSamurai())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if globalErr != nil {
		log.Fatal(globalErr)
	}
}
