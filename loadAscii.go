package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type loadAsciiMsg struct {
	lg  []string
	md  []string
	sm  []string
	err error
}

func loadAscii() tea.Msg {
	return loadAsciiMsg{
		lg: samuraiLarge,
		md: samuraiMedium,
		sm: samuraiSmall,
	}
}
