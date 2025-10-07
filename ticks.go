package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const MAX_FPS = 15

type bannerTick struct{}

func (m model) tickBanner() tea.Cmd {
	fps := time.Duration(1000 / MAX_FPS)

	return tea.Tick(time.Millisecond*fps, func(t time.Time) tea.Msg {
		return bannerTick{}
	})
}
