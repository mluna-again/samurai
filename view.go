package main

import "github.com/charmbracelet/lipgloss"

func (m model) banner() string {
	if !m.ready {
		return "Loading..."
	}

	banner := m.currentBanner[m.currentFrame]
	if m.layout == VERTICAL {
		banner = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, banner)
	}

	return banner
}
