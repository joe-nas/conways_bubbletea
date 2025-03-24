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
	name   string // header
	index  int    // 0
	height int
	style  lipgloss.Style
}

func (r Row) String() string {
	return r.name
}

type Column struct {
	name  string // header
	index int
	width int
	style lipgloss.Style
}

type Cell struct {
	name    string
	content string
	rowIdx  int
	colIdx  int
	style   lipgloss.Style
}

func (g Grid) addCell(rowIdx, colIdx int, cell *Cell) {
	g.layout[rowIdx][colIdx].append(cell)
}

// func (gl *Grid) newGridLayout(nrows, ncols int) {
// 	gl.layout := make([]Row, nrows)
// }

func newGrid() *Grid {
	return &Grid{
		layout:  nil,
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

func main() {

	mapstring :=
		`sldkf;lskdf
sdf;lskdf;
;lskdf;
';lsd'f;lsd'f;l'
sdplf[psf[plsdfls';dlf'
sd;f'lksd;lks;dlfk;sldkf]]`

	gamemap_cell := Cell{
		name:    "game map",
		content: mapstring,
		rowIdx:  0,
		colIdx:  0,
	}

	headerRow := Row{name: "header", index: 0, height: 10, style: headerStyle}
	bodyRow := Row{name: "header", index: 0, height: 10, style: bodyStyle}
	footerRow := Row{name: "footer", index: 2, height: 10, style: footerStyle}
	grid := newGrid()
	grid.layout[0] = new(Cell)
	grid.addRow(headerRow, bodyRow, footerRow)
	grid.addColumn(Column{name: "Stats left", index: 0, width: 10}, Column{name: "Stats right", index: 1, width: 10})
	grid.addCell(gamemap_cell)
	fmt.Print(grid)
	// fmt.Print(headerRow)

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
	headerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7d33ff")).
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
