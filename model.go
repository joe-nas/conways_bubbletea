package main

import (
	"conways_bubbletea/styles"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Row metadata
	titleRow  = row{name: "title", index: 0, style: styles.TitleStyle}
	headerRow = row{name: "header", index: 1, style: styles.HeaderStyle}
	helpRow   = row{name: "header", index: 2, style: styles.HeaderStyle}
	bodyRow   = row{name: "header", index: 3, style: styles.BodyStyle}
	footerRow = row{name: "footer", index: 4, style: styles.FooterStyle}
)

var (
	//  Column metadata
	statLeft   = column{name: "Stats left", index: 0, style: styles.HeaderStyle}
	statCenter = column{name: "Stats center", index: 1, style: styles.TitleStyle}
	statRight  = column{name: "Stats right", index: 2, style: styles.FooterStyle}
)

func (g grid) renderRowContent(rowIdx int) string {
	var s string
	fmt.Println(len(g.layout[rowIdx]))
	for colIdx := range g.layout[rowIdx] {
		s = lipgloss.JoinHorizontal(lipgloss.Top, s, g.columns[colIdx].style.Render(g.layout[rowIdx][colIdx].content))
	}
	return s
}

type tile struct {
	curr_gen  bool
	next_gen  bool
	neighbors int
}

type grid struct {
	layout  [][]cell // actual data
	rows    []row    // store metadata
	columns []column // store metadata
}

type row struct {
	name  string // header
	index int    // 0
	// height int
	style lipgloss.Style
}

type column struct {
	name  string // header
	index int
	// width int
	style lipgloss.Style
}

type cell struct {
	name    string
	content string
	rowIdx  int
	colIdx  int
	style   lipgloss.Style
}

func newGrid(cols ...int) *grid {
	// cols parameter represents the number of colums per Rows...
	// newGrid(1,2,3) creates a Layout with 1 column in row 0, 2 Columns in row 1 and 3 Columns in row 2
	layout := make([][]cell, len(cols))
	fmt.Println("---------")
	fmt.Println(layout)
	fmt.Println("---------")
	for row := range len(layout) {
		layout[row] = make([]cell, cols[row])
	}

	return &grid{
		layout:  layout,
		rows:    nil,
		columns: nil,
	}
}
func (g grid) addCell(cell *cell) {
	g.layout[cell.rowIdx][cell.colIdx] = *cell
}

func (g *grid) addRow(row ...row) {
	g.rows = append(g.rows, row...)
}

func (g *grid) addColumn(column ...column) {
	g.columns = append(g.columns, column...)
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

func renderCell(cell *cell) string {
	return cell.style.Render(cell.content)
}

func (m model) renderHeader() string {

	// header := "Conway's Game Of Life\n"
	// header += fmt.Sprintf("cursor pos: x:%d y:%d \t\t Autorun: %t\n", m.cursor["x"], m.cursor["y"], m.autorun)
	// header += fmt.Sprintf("Current generation: %d\n", m.generation)
	// header += fmt.Sprintf("Map Size: x:%d y:%d\n", len(m.matrix), len(m.matrix[0]))
	// header += "(e) toggle tile state, (n) activate next generation\n"
	// header += "(a) toggle between map and neighbor count view, (q) quit"
	// return headerStyle.Render(header)
	header := renderCell(&titleCell)
	return header
}

func (m model) Init() tea.Cmd {
	return nil
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
	return styles.BodyStyle.Render(game_map)
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

func (m model) View() string {
	header := m.renderHeader()

	var content string
	if !m.altscreen {
		content = m.renderGameMap()
	} else {
		content = m.renderNeighborCount()
	}
	return header + content + "\n"
}
