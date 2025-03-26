package main

import (
	"conways_bubbletea/styles"
)

var (
	// row 0
	titleCell = cell{name: "Title", content: "Conway's Game of Life", rowIdx: 0, colIdx: 0, style: styles.TitleStyle}
	// row 1
	statLeftCell   = cell{name: "Stat left", content: "Cursor pos: x:%d y:%d\nMap Size: x:%d y:%d", rowIdx: 1, colIdx: 0, style: styles.HeaderStyle}
	statCenterCell = cell{name: "Stat center", content: "Autorun: %t", rowIdx: 1, colIdx: 1, style: styles.HeaderStyle}
	statRightCell  = cell{name: "Stat right", content: "Current generation: %d", rowIdx: 1, colIdx: 2, style: styles.HeaderStyle}
	// row 2
	helpCell = cell{name: "Help cell", content: `(e) toggle tile state, (n) activate next generation
	(a) toggle between map and neighbor count view, (q) quit`, rowIdx: 2, colIdx: 0, style: styles.HelpStyle}
)
