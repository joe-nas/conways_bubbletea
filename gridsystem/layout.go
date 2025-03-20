package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Layout struct {
	layout_matrix [][]Column
}

type Column struct {
	name    string
	content string
	style   lipgloss.Style
	// col int
}

// func (l Layout) rowRender() string {
// 	l.layout_matrix[]
// }

// func (c Column) newColumn(width, height int, name string) {
// 	return
// }

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
			Background(lipgloss.Color("#7d33ff")).
			Width(10).
			Height(10).BorderStyle(lipgloss.RoundedBorder())

	teststyle2 = lipgloss.NewStyle().
			Background(lipgloss.Color("#ff5733")).
			Width(10).
			Height(7).BorderStyle(lipgloss.RoundedBorder())
)

func (c *Column) addContentToColumn(name string, s string) {
	c.name = name
	c.content = teststyle.Render(s)
}

func main() {
	var layout = NewLayout()
	// row0 - 2

	layout.addColumns(1)
	layout.layout_matrix[0] = []Column{{name: "header", content: "jernvnb,"}}
	layout.addColumns(2)
	layout.layout_matrix[1] = []Column{{name: "stats", content: teststyle.Render("fsdfsdf;k ;dkf;gkd;fg ;ldfg;ldkfg  0e-r0ter")}, {name: "stats", content: teststyle2.Render("hoeg")}}
	layout.addColumns(10)
	// layout.layout_matrix[2][3].addContentToColumn(teststyle.Render("0000345384598"))

	// fmt.Print(layout)
	lm := &layout.layout_matrix
	for i := range *lm {
		for j := range (*lm)[i] {
			lipgloss.JoinHorizontal(lipgloss.Bottom, (*lm)[i][j].content)
		}
	}

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, (*lm)[1][0].content, (*lm)[1][1].content, (*lm)[1][0].content, (*lm)[1][1].content))
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
