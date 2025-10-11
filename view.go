package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) banner() string {
	if !m.ready {
		return "Loading..."
	}

	banner := m.currentBanner[m.currentFrame]

	return banner
}

func (m model) sessionList() string {
	width := lipgloss.Width(m.banner())
	if width > (m.width / 2) {
		if m.layout == VERTICAL {
			width = m.width
		} else {
			width = m.width - width
		}
	}

	lines := []string{lipgloss.PlaceHorizontal(width, lipgloss.Center, sessionTitle)}
	margin := 4
	for _, sess := range m.sessions {
		right := sess[0]
		left := sess[1]
		w := width - lipgloss.Width(right) - lipgloss.Width(left) - margin*2
		if w < 0 {
			w = 0
		}

		mar := strings.Repeat(" ", margin)
		padd := strings.Repeat(" ", w)
		line := lipgloss.JoinHorizontal(lipgloss.Center, mar, right, padd, left, mar)

		lines = append(lines, line)
	}

	return lipgloss.JoinVertical(lipgloss.Center, lines...)
}
