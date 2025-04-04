package main

import "conways_bubbletea/styles"

var (
	// row 0
	titleCell = cell{
		name:    "title",
		content: "Conway's Game of Life",
		rowIdx:  0,
		colIdx:  0,
		style:   styles.TitleStyle,
	}
	// row 1
	statLeftCell = cell{
		name:    "statLeft",
		content: "Cursor pos: x:%d y:%d\nMap Size: x:%d y:%d",
		rowIdx:  1,
		colIdx:  0,
		style:   styles.StatsStyle,
	}
	statCenterCell = cell{
		name:    "statCenter",
		content: "Alive Cells: %d\nDead Cells: %d",
		rowIdx:  1,
		colIdx:  1,
		style:   styles.StatsStyle,
	}
	statRightCell = cell{
		name:    "statRight",
		content: "Autorun: %t\nCurrent generation: %d",
		rowIdx:  1,
		colIdx:  2,
		style:   styles.StatsStyle,
	}
	// row 2
	helpCell = cell{
		name: "helpCell",
		content: "(e) toggle tile state, (n) activate next generation, (a) toggle autorun\n" +
			"(r) reset game state, (s) toggle between map and neighbor count view, (q) quit",
		rowIdx: 2,
		colIdx: 0,
		style:  styles.HelpStyle,
	}

	// row 4
	countViewHelp = cell{
		name: "countExplainer",
		content: "Alive: %d, Surviving: %d, Born: %d, \n" +
			"Dying: %d, Dead %d",
		rowIdx: 4,
		colIdx: 0,
		style:  styles.HelpStyle,
	}
)

func (m *model) addCell(c cell) {
	m.layout.layout[c.rowIdx][c.colIdx] = c
}
