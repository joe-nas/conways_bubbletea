package main

type layout struct {
	layout_matrix [][]string
}
type column struct {
	s   string
	row int
	col int
}

func (l layout) addColumn() {
	l.layout_matrix = append(l.layout_matrix)
}

func (l layout) addRow() {

}

func createLayout() layout {

	header := []string{"header"}
	stats := []string{"stat1", "stat2", "stat3"}
	help := []string{"help"}
	game := []string{"game"}
	footer := []string{"footer"}

	return layout{
		layout_matrix: [][]string{header, stats, help, game, footer},
	}
}
