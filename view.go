package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true)

	// Background(lipgloss.Color("#FFFF00")). // Yellow background
	// Padding(0, 1)

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7d33ff"))

	deadStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#000000"))

	dieStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FF0000"))

	birthStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90ee90"))

	surviveStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90eeee"))

	// modalStyle = lipgloss.NewStyle().

)

func (m model) RenderHeader() string {

	header := "Conway's Game Of Life\n"
	header += fmt.Sprintf("Cursor pos: x:%d y:%d\n", m.cursor["x"], m.cursor["y"])
	header += fmt.Sprintf("Current generation: %d\n", m.generation)
	header += fmt.Sprintf("Map Size: x:%d y:%d\n", len(m.matrix), len(m.matrix[0]))
	header += "(e) toggle tile state, (n) activate next generation\n"
	header += "(a) toggle between map and neighbor count view, (q) quit"
	return headerStyle.Render(header)
}

func (m model) RenderGameMap() string {
	var game_map string
	for i := range m.matrix {
		game_map += "\n"
		for j := range m.matrix[i] {
			var cell string
			if m.matrix[i][j].curr_gen {
				cell = "👽"
				// cell = "o"
			} else {
				cell = "💀"
				// cell = "x"
			}

			// cursor highlighter
			if i == m.cursor["x"] && j == m.cursor["y"] {
				cell = cursorStyle.Render(cell)
			}
			game_map += cell
		}
	}
	return headerStyle.Render(game_map)
}

func (m model) RenderModal() string {
	// lipgloss.Place(10, 10, lipgloss.Center, lipgloss.Center, headerStyle.Render("Dies ist ein test"))
	return lipgloss.Place(10, 10, lipgloss.Right, lipgloss.Top, headerStyle.Render("Dies ist ein test"))
}

func (m model) RenderNeighborCount() string {
	var neighbor_count string
	var styled_cell string
	for i := range m.matrix {
		neighbor_count += "\n"
		for j := range m.matrix[i] {

			cell := &m.matrix[i][j]
			cellText := fmt.Sprintf(" %d", cell.neighbors)

			if cell.curr_gen {
				// Currently alive cell
				if cell.neighbors < 2 {
					// Death by isolation: Fewer than 2 neighbors
					styled_cell = dieStyle.Render(cellText)
				} else if cell.neighbors == 2 || cell.neighbors == 3 {
					// Survival: 2 or 3 neighbors
					styled_cell = surviveStyle.Render(cellText)
				} else {
					// Death by overcrowding: More than 3 neighbors
					styled_cell = dieStyle.Render(cellText)
				}
			} else {
				// Currently dead cell
				if cell.neighbors == 3 {
					// Birth: Exactly 3 neighbors
					styled_cell = birthStyle.Render(cellText)
				} else {
					// Remain dead
					styled_cell = deadStyle.Render(cellText)
				}
			}

			// render cursor
			if i == m.cursor["x"] && j == m.cursor["y"] {
				styled_cell = cursorStyle.Render(cellText)
			}
			neighbor_count += styled_cell
		}
	}
	return neighbor_count
}

func (m model) View() string {
	header := m.RenderHeader()

	var content string
	if !m.altscreen {
		content = m.RenderGameMap()
	} else {
		content = m.RenderNeighborCount()
	}
	return header + content + "\n"
}
