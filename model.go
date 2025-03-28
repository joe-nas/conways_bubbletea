package main

import (
	"conways_bubbletea/styles"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// var (
// 	// Row metadata
// 	titleRow  = row{name: "title", index: 0, style: styles.TitleStyle}
// 	headerRow = row{name: "header", index: 1, style: styles.HeaderStyle}
// 	helpRow   = row{name: "header", index: 2, style: styles.HeaderStyle}
// 	bodyRow   = row{name: "header", index: 3, style: styles.BodyStyle}
// 	footerRow = row{name: "footer", index: 4, style: styles.FooterStyle}
// )

// var (
// 	//  Column metadata
// 	statLeft   = column{name: "Stats left", index: 0, style: styles.HeaderStyle}
// 	statCenter = column{name: "Stats center", index: 1, style: styles.TitleStyle}
// 	statRight  = column{name: "Stats right", index: 2, style: styles.FooterStyle}
// )

type tile struct {
	curr_gen  bool
	next_gen  bool
	neighbors int
}

type model struct {
	layout grid
	// state
	matrix     [][]tile
	cursor     map[string]int
	altscreen  bool
	autorun    bool
	nrows      int
	ncols      int
	generation int
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel(nrows int, ncols int) model {
	layout := newGrid(1, 3, 1, 2, 1) // title, stats, help, game, footer

	matrix := make([][]tile, nrows)
	for row := range matrix {
		matrix[row] = make([]tile, ncols)
	}

	return model{
		layout:     *layout,
		matrix:     matrix,
		cursor:     map[string]int{"x": 0, "y": 0},
		nrows:      nrows,
		ncols:      ncols,
		generation: 0,
		autorun:    false,
	}
}

func (c cell) renderCell(m model) string {
	switch c.name {
	case "title":
		return c.style.Render(c.content)
	case "statLeft":
		return c.style.Render(fmt.Sprintf(c.content, m.cursor["x"], m.cursor["y"], m.nrows, m.ncols))
	case "statCenter":
		return c.style.Render(fmt.Sprintf(c.content, m.autorun))
	case "statRight":
		return c.style.Render(fmt.Sprintf(c.content, m.generation))
	default:
		return c.style.Render(c.content)
	}
}

func (m model) renderGameMap() string {
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
				cell = styles.CursorStyle.Render(cell)
			}
			game_map += cell
		}
	}
	return game_map
}

func (m model) renderNeighborCount() string {
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
					styled_cell = styles.DieStyle.Render(cellText)
				} else if cell.neighbors == 2 || cell.neighbors == 3 {
					// Survival: 2 or 3 neighbors
					styled_cell = styles.SurviveStyle.Render(cellText)
				} else {
					// Death by overcrowding: More than 3 neighbors
					styled_cell = styles.DieStyle.Render(cellText)
				}
			} else {
				// Currently dead cell
				if cell.neighbors == 3 {
					// Birth: Exactly 3 neighbors
					styled_cell = styles.BirthStyle.Render(cellText)
				} else {
					// Remain dead
					styled_cell = styles.DeadStyle.Render(cellText)
				}
			}

			// render cursor
			if i == m.cursor["x"] && j == m.cursor["y"] {
				styled_cell = styles.CursorStyle.Render(cellText)
			}
			neighbor_count += styled_cell
		}
	}
	return neighbor_count
}
