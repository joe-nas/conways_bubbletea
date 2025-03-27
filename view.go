package main

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	header := renderCell(&titleCell)

	var content string
	if !m.altscreen {
		content = m.renderGameMap()
	} else {
		content = m.renderNeighborCount()
	}
	return lipgloss.JoinVertical(lipgloss.Left, header, content)
}
