package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// grid represents the overall Layout of the tui
// layout stores the cells data as well as row and column metadata
type grid struct {
	layout  [][]cell // actual data
	rows    []row    // store metadata
	columns []column // store metadata
}

// cell represents a cell within the grid layout
// it holds the actual data of a cell as well as style data and position
type cell struct {
	name    string
	content string
	rowIdx  int
	colIdx  int
	style   lipgloss.Style
}

// row holds metadata about a particular row in the grid
type row struct {
	name  string // header
	index int    // 0
	// height int
	style lipgloss.Style
}

// column holds metadata about a particular column in the grid
type column struct {
	name  string // header
	index int
	// width int
	style lipgloss.Style
}

// newGrid returns a pointer to a grid
// cols parameter represents the number of colums per Rows.
// newGrid(1,2,3) creates a Layout with 1 column in row 0, 2 Columns in row 1 and 3 Columns in row 2
func newGrid(cols ...int) *grid {
	layout := make([][]cell, len(cols))
	for row := range len(layout) {
		layout[row] = make([]cell, cols[row])
	}

	return &grid{
		layout:  layout,
		rows:    nil,
		columns: nil,
	}
}

// addCell can add a cell to the grid
func (g grid) addCell(cell *cell) {
	g.layout[cell.rowIdx][cell.colIdx] = *cell
}

// renderRowContent renders and joins the columns of a particular row (rowIdx) of the grid
func (g grid) renderRowContent(rowIdx int) string {
	var s string
	fmt.Println(len(g.layout[rowIdx]))
	for colIdx := range g.layout[rowIdx] {
		s = lipgloss.JoinHorizontal(lipgloss.Top, s, g.columns[colIdx].style.Render(g.layout[rowIdx][colIdx].content))
	}
	return s
}

// addMetadataRow can add a Metadata row to the grid object
func (g *grid) addMetadataRow(row ...row) {
	g.rows = append(g.rows, row...)
}

// addColumndataRow can add a Metadata column to the grid object
func (g *grid) addMetadataColumn(column ...column) {
	g.columns = append(g.columns, column...)
}
