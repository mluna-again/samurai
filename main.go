package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// FLAGS
var noanimate bool
var separator string
var sessionTitle string
var rightDates string

// FLAGS

var response string = ""
var forcedLayout string

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
	sessions      [][]string
}

func newSamurai() (model, error) {
	sessions, err := loadSessions()
	if err != nil {
		return model{}, err
	}

	l := VERTICAL
	if forcedLayout == "vertical" {
		l = VERTICAL
	} else if forcedLayout == "horizontal" {
		l = HORIZONTAL
	}

	return model{
		layout:   l,
		sessions: sessions,
	}, nil
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
		if forcedLayout == "" {
			if lipgloss.Height(m.View()) > m.heigth {
				m.layout = HORIZONTAL
			} else {
				m.layout = VERTICAL
			}
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
		keystroke := msg.String()
		if len(m.sessions) > 0 {
			num, err := strconv.Atoi(keystroke)
			if err == nil && num >= 0 && num < len(m.sessions) {
				response = m.sessions[num][0]
				return m, tea.Quit
			}
		}

		switch keystroke {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	banner := m.banner()
	if m.layout == VERTICAL {
		banner = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, banner)
	}

	list := m.sessionList()
	if m.layout == VERTICAL {
		list = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, list)
		return lipgloss.JoinVertical(lipgloss.Top, banner, list)
	}

	list = lipgloss.PlaceVertical(m.heigth, lipgloss.Center, list)
	content := lipgloss.JoinHorizontal(lipgloss.Center, banner, list)

	return lipgloss.Place(m.width, m.heigth, lipgloss.Center, lipgloss.Center, content)
}

func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)

	flag.BoolVar(&noanimate, "noanimate", false, "don't animate ascii art")
	flag.StringVar(&separator, "sep", "@", "component separator")
	flag.StringVar(&sessionTitle, "title", " Recent sessions ", "Sessions header")
	flag.StringVar(&rightDates, "rformat", time.Kitchen, "Treats right components as epoch values and formats them with the given format, an empty string disables this")
	flag.StringVar(&forcedLayout, "layout", "", "layout. by default it automatically changes when the screen is resized to try to fit it, but you can force a specific one. can be either vertical or horizontal.")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage\n")
		fmt.Fprintf(os.Stderr, "samurai reads from stdin and accepts up to 10 lines. each line can have 2 components, one on the right and one on the left, they should be separated by -sep\n\n")
		fmt.Fprintf(os.Stderr, "$ tmux list-sessions -a -F '#{session_name}@#{session_created}' | samurai -sep @\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	m, err := newSamurai()
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(m, tea.WithOutput(os.Stderr), tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if globalErr != nil {
		log.Fatal(globalErr)
	}

	if response != "" {
		fmt.Println(response)
	}
}
