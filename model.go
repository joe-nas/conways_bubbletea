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

type deadOrAliveStats struct {
	dead  int
	alive int
}

type neighborViewStats struct {
	dead int
	born int
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

	neighborViewStats *neighborViewStats
	deadOrAliveStats  *deadOrAliveStats
}

func (m model) Init() tea.Cmd {
	return nil
}

func initialModel(nrows int, ncols int) model {
	layout := newGrid(1, 3, 1, 2, 1) // title, stats, help, game, footer

	deadOrAliveStats := deadOrAliveStats{0, 0}

	matrix := make([][]tile, nrows)
	for row := range matrix {
		matrix[row] = make([]tile, ncols)
	}

	return model{
		layout:           *layout,
		matrix:           matrix,
		cursor:           map[string]int{"x": 0, "y": 0},
		nrows:            nrows,
		ncols:            ncols,
		generation:       0,
		autorun:          false,
		deadOrAliveStats: &deadOrAliveStats,
	}
}

func (c cell) renderCell(m model) string {
	switch c.name {
	case "title":
		return c.style.Render(c.content)
	case "statLeft":
		return c.style.Render(fmt.Sprintf(c.content, m.cursor["x"], m.cursor["y"], m.nrows, m.ncols))
	case "statCenter":
		return c.style.Render(fmt.Sprintf(c.content, m.deadOrAliveStats.alive, m.deadOrAliveStats.dead))
	case "statRight":
		return c.style.Render(fmt.Sprintf(c.content, m.autorun, m.generation))
	default:
		return c.style.Render(c.content)
	}
}

func (m *model) countDeadAlive() {
	m.deadOrAliveStats.alive = 0
	m.deadOrAliveStats.dead = 0
	for i := range m.matrix {
		for j := range m.matrix[i] {
			if m.matrix[i][j].curr_gen {
				m.deadOrAliveStats.alive += 1
			} else {
				m.deadOrAliveStats.dead += 1
			}
		}
	}
}

func (m *model) renderGameMap() string {
	var game_map string
	// m.neighorStats.alive = 0
	// m.neighorStats.dead = 0
	for i := range m.matrix {
		game_map += "\n"
		for j := range m.matrix[i] {
			var cell string
			if m.matrix[i][j].curr_gen {
				// m.neighorStats.alive += 1
				cell = "👽"
				// cell = "o"
			} else {
				// m.neighorStats.dead += 1
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
	m.countDeadAlive()
	return game_map
}

func (m *model) renderNeighborCount() string {
	var neighbor_count string
	var styled_cell string
	// m.neighorStats.alive = 0
	// m.neighorStats.dead = 0
	for i := range m.matrix {
		neighbor_count += "\n"
		for j := range m.matrix[i] {

			cell := &m.matrix[i][j]
			cellText := fmt.Sprintf(" %d", cell.neighbors)

			if cell.curr_gen {
				// Currently alive cell
				if cell.neighbors < 2 {
					// Death by isolation: Fewer than 2 neighbors
					// m.neighorStats.dead += 1
					styled_cell = styles.CellDieingStyle.Render(cellText)
				} else if cell.neighbors == 2 || cell.neighbors == 3 {
					// Survival: 2 or 3 neighbors
					// m.neighorStats.alive += 1
					styled_cell = styles.CellSurviveStyle.Render(cellText)
				} else {
					// Death by overcrowding: More than 3 neighbors
					// m.neighorStats.dead += 1
					styled_cell = styles.CellDieingStyle.Render(cellText)
				}
			} else {
				// Currently dead cell
				if cell.neighbors == 3 {
					// Birth: Exactly 3 neighbors
					// m.neighorStats.alive += 1
					styled_cell = styles.CellBirthStyle.Render(cellText)
				} else {
					// Remain dead
					// m.neighorStats.dead += 1
					styled_cell = styles.CellDeadStyle.Render(cellText)
				}
			}

			// render cursor
			if i == m.cursor["x"] && j == m.cursor["y"] {
				styled_cell = styles.CursorStyle.Render(cellText)
			}
			neighbor_count += styled_cell
		}
	}
	m.countDeadAlive()
	return neighbor_count
}
