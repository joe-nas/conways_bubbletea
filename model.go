package main

type Tile struct {
	curr_gen  bool
	next_gen  bool
	neighbors int
}

type GameState struct {
	Matrix   [][]Tile
	Cursor   map[string]int
	ViewMode string
}
