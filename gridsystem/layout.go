package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Layout struct {
	layout_matrix [][]Column
}

type Column struct {
	s string
	// row int
	// col int
}

func NewLayout() *Layout {
	layout := new(Layout)
	// layout.layout_matrix = make([][]Column, 0)
	return layout
}

func (l *Layout) addColumns(columns int) {
	l.layout_matrix = append(l.layout_matrix, make([]Column, columns))
}

var (
	teststyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#7d33ff"))
)

func (c *Column) addContentToColumn(s string) {
	c.s = teststyle.Render(s)
}

func main() {
	var layout = NewLayout()
	// row0 - 2

	layout.addColumns(1)
	layout.layout_matrix[0] = []Column{{s: "jernvnb,"}}
	layout.addColumns(2)
	layout.layout_matrix[1] = []Column{{s: teststyle.Render("fsdfsdf")}, {s: "hoeg"}}
	layout.addColumns(10)
	layout.layout_matrix[2][3].addContentToColumn(teststyle.Render("0000345384598"))

	// fmt.Print(layout)
	lm := &layout.layout_matrix
	for i := range *lm {
		for j := range (*lm)[i] {
			fmt.Print((*lm)[i][j])
		}
		fmt.Println()

	}
}

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
