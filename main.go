package main

import (
	"flag"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var noanimate bool

const SM_BREAK = 50
const MD_BREAK = 60

type layout int

const (
	VERTICAL layout = iota
	HORIZONTAL
)

var globalErr error = nil

type model struct {
	smBanner      []string
	mdBanner      []string
	lgBanner      []string
	currentBanner []string
	currentFrame  int
	layout        layout
	ready         bool
	width         int
	heigth        int
}

func newSamurai() model {
	return model{
		layout: VERTICAL,
	}
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
		if noanimate {
			return m, nil
		}
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
		if lipgloss.Height(m.View()) > m.heigth {
			m.layout = HORIZONTAL
		} else {
			m.layout = VERTICAL
		}
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
	flag.BoolVar(&noanimate, "noanimate", false, "don't animate ascii art")
	flag.Parse()

	p := tea.NewProgram(newSamurai())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if globalErr != nil {
		log.Fatal(globalErr)
	}
}
