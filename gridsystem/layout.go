package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Grid struct {
	layout  [][]Cell // actual data
	rows    []Row    // store metadata
	columns []Column // store metadata
}

type Row struct {
	name  string // header
	index int    // 0
	// height int
	style lipgloss.Style
}

func (r Row) String() string {
	return r.name
}

type Column struct {
	name  string // header
	index int
	// width int
	style lipgloss.Style
}

type Cell struct {
	name    string
	content string
	rowIdx  int
	colIdx  int
	style   lipgloss.Style
}

func (g Grid) addCell(cell *Cell) {
	g.layout[cell.rowIdx][cell.colIdx] = *cell
}

// func (gl *Grid) newGridLayout(nrows, ncols int) {
// 	gl.layout := make([]Row, nrows)
// }

// func newGrid(nrows, ncols int) *Grid {
// 	layout := make([][]Cell, nrows)
// 	for row := range layout {
// 		layout[row] = make([]Cell, ncols)
// 	}

// 	return &Grid{
// 		layout:  layout,
// 		rows:    nil,
// 		columns: nil,
// 	}
// }

func newGrid(cols ...int) *Grid {
	// cols parameter represents the number of colums per rows...
	// newGrid(1,2,3) creates a layout with 1 column in row 0, 2 columns in row 1 and 3 columns in row 2
	layout := make([][]Cell, len(cols))
	fmt.Println("---------")
	fmt.Println(layout)
	fmt.Println("---------")
	for row := range len(layout) {
		layout[row] = make([]Cell, cols[row])
	}

	return &Grid{
		layout:  layout,
		rows:    nil,
		columns: nil,
	}
}

func (g *Grid) addRow(row ...Row) {
	g.rows = append(g.rows, row...)
}

func (g *Grid) addColumn(column ...Column) {
	g.columns = append(g.columns, column...)
}

var (
	// row 0
	titleCell = Cell{name: "Title", content: "Conway's Game of Life", rowIdx: 0, colIdx: 0, style: titleStyle}
	// row 1
	statLeftCell   = Cell{name: "Stat left", content: "Cursor pos: x:%d y:%d\nMap Size: x:%d y:%d", rowIdx: 1, colIdx: 0, style: headerStyle}
	statCenterCell = Cell{name: "Stat center", content: "Autorun: %t", rowIdx: 1, colIdx: 1, style: headerStyle}
	statRightCell  = Cell{name: "Stat right", content: "Current generation: %d", rowIdx: 1, colIdx: 2, style: headerStyle}
	// row 2
	helpCell = Cell{name: "Help cell", content: `(e) toggle tile state, (n) activate next generation
	(a) toggle between map and neighbor count view, (q) quit`, rowIdx: 2, colIdx: 0, style: helpStyle}

// header += fmt.Sprintf("Cursor pos: x:%d y:%d \t\t Autorun: %t\n", m.cursor["x"], m.cursor["y"], m.autorun)
// header += fmt.Sprintf("Current generation: %d\n", m.generation)
// header += fmt.Sprintf("Map Size: x:%d y:%d\n", len(m.matrix), len(m.matrix[0]))
// header += "(e) toggle tile state, (n) activate next generation\n"
// header += "(a) toggle between map and neighbor count view, (q) quit")
)

func (g Grid) renderRowContent(rowIdx int) string {
	var s string
	fmt.Println(len(g.layout[rowIdx]))
	for colIdx := range g.layout[rowIdx] {
		s = lipgloss.JoinHorizontal(lipgloss.Top, s, g.columns[colIdx].style.Render(g.layout[rowIdx][colIdx].content))
	}
	return s
}

func main() {

	// 	mapstring :=
	// 		`sldkf;lskdf
	// sdf;lskdf;
	// ;lskdf;
	// ';lsd'f;lsd'f;l'
	// sdplf[psf[plsdfls';dlf'
	// sd;f'lksd;lks;dlfk;sldkf]]`

	// gamemap_cell := Cell{
	// 	name:    "game map",
	// 	content: mapstring,
	// 	rowIdx:  0,
	// 	colIdx:  0,
	// }

	// Row metadata
	titleRow := Row{name: "title", index: 0, style: titleStyle}
	headerRow := Row{name: "header", index: 1, style: headerStyle}
	helpRow := Row{name: "header", index: 2, style: headerStyle}
	bodyRow := Row{name: "header", index: 3, style: bodyStyle}
	footerRow := Row{name: "footer", index: 4, style: footerStyle}

	//  Column metadata
	statLeft := Column{name: "Stats left", index: 0, style: headerStyle}
	statCenter := Column{name: "Stats center", index: 1, style: titleStyle}
	statRight := Column{name: "Stats right", index: 2, style: footerStyle}

	grid := newGrid(1, 3, 2)
	fmt.Println(grid.layout)
	grid.addColumn(statLeft, statCenter, statRight)
	grid.addCell(&titleCell)
	grid.addRow(titleRow, headerRow, helpRow, bodyRow, footerRow)
	grid.addCell(&statLeftCell)
	grid.addCell(&statCenterCell)
	grid.addCell(&statRightCell)
	grid.addCell(&helpCell)

	titlecellstring := grid.renderRowContent(0)
	headerrowstring := grid.renderRowContent(1)
	statsrowstring := grid.renderRowContent(2)
	fmt.Println(lipgloss.JoinVertical(lipgloss.Center, titlecellstring, headerrowstring, statsrowstring))
	fmt.Println(grid.layout[0])
}

// func (l Layout) rowRender() string {
// 	l.layout_matrix[]
// }

// func (c Column) newColumn(width, height int, name string) {
// 	return
// }

// func NewLayout() *Layout {
// 	layout := new(Layout)
// 	// layout.layout_matrix = make([][]Column, 0)
// 	return layout
// }

// func (l *Layout) addColumns(columns int) {
// 	l.layout_matrix = append(l.layout_matrix, make([]Column, columns))
// }

var (
	titleStyle = lipgloss.NewStyle().
			Underline(true).
			Height(2)
		// Background(lipgloss.Color("#7d33ff")).
		// Width(10).
		// Height(10).
		// BorderStyle(lipgloss.ThickBorder())

	headerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7d33ff")).
			Width(10).
			Height(10).BorderStyle(lipgloss.RoundedBorder())

	helpStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#88d4c3")).
			Width(10).
			Height(10).BorderStyle(lipgloss.RoundedBorder())

	bodyStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#32cd32")).
			Width(10).
			Height(7).BorderStyle(lipgloss.RoundedBorder())

	footerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#ff5733")).
			Width(10).
			Height(7).BorderStyle(lipgloss.RoundedBorder())
)

// func (c *Column) addContentToColumn(name string, s string) {
// 	c.name = name
// 	c.content = teststyle.Render(s)
// }

// func main() {
// 	var layout = NewLayout()
// 	// row0 - 2

// 	layout.addColumns(1)
// 	layout.layout_matrix[0] = []Column{{name: "header", content: "jernvnb,"}}
// 	layout.addColumns(2)
// 	layout.layout_matrix[1] = []Column{{name: "stats", content: teststyle.Render("fsdfsdf;k ;dkf;gkd;fg ;ldfg;ldkfg  0e-r0ter")}, {name: "stats", content: teststyle2.Render("hoeg")}}
// 	layout.addColumns(10)
// 	// layout.layout_matrix[2][3].addContentToColumn(teststyle.Render("0000345384598"))

// 	// fmt.Print(layout)
// 	lm := &layout.layout_matrix
// 	for i := range *lm {
// 		for j := range (*lm)[i] {
// 			lipgloss.JoinHorizontal(lipgloss.Bottom, (*lm)[i][j].content)
// 		}
// 	}

// 	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, (*lm)[1][0].content, (*lm)[1][1].content, (*lm)[1][0].content, (*lm)[1][1].content))
// }

// func (l layout) addColumn() {
// 	// l.layout_matrix = append(l.layout_matrix)
// }

// func layoutCreate() *layout {
// 	return &layout{}
// }

// func createLayout() layout {

// 	header := []string{"header"}
// 	stats := []string{"stat1", "stat2", "stat3"}
// 	help := []string{"help"}
// 	game := []string{"game"}
// 	footer := []string{"footer"}

// 	return layout{
// 		layout_matrix: [][]string{header, stats, help, game, footer},
// 	}
// }
