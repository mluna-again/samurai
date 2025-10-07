package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

const SM_BREAK = 50
const MD_BREAK = 60

var globalErr error = nil

type model struct {
	smBanner      []string
	mdBanner      []string
	lgBanner      []string
	currentBanner []string
	currentFrame  int
	ready         bool
	width         int
	heigth        int
}

func newSamurai() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return loadAscii
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case bannerTick:
		m.currentFrame++
		if m.currentFrame >= len(m.currentBanner) {
			m.currentFrame = 0
		}
		return m, m.tickBanner()

	case loadAsciiMsg:
		if msg.err != nil {
			globalErr = msg.err
			return m, tea.Quit
		}
		m.smBanner = msg.sm
		m.mdBanner = msg.md
		m.lgBanner = msg.lg
		m.currentBanner = m.lgBanner
		m.ready = true
		return m, m.tickBanner()

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.heigth = msg.Height
		if m.width < SM_BREAK {
			m.currentBanner = m.smBanner
		} else if m.width < MD_BREAK {
			m.currentBanner = m.mdBanner
		} else {
			m.currentBanner = m.lgBanner
		}
		m.currentFrame = 0
		return m, tea.ClearScreen

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.banner()
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
