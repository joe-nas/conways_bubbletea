package main

import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
	title := titleCell.renderCell(m)

	statsLeft := statLeftCell.renderCell(m)
	statsCenter := statCenterCell.renderCell(m)
	statsRight := statRightCell.renderCell(m)
	stats := lipgloss.JoinHorizontal(lipgloss.Bottom, statsLeft, statsCenter, statsRight)

	help := helpCell.renderCell(m)

	var content string
	if !m.altscreen {
		content = m.renderGameMap()
	} else {
		content = m.renderNeighborCount()
	}
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(lipgloss.JoinVertical(lipgloss.Center, title, stats, help, content))
}
