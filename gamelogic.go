package main

func (m model) countNeighbors() {
	var neighbors = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var nx_cord, ny_cord int

	nrows := len(m.matrix)
	ncols := len(m.matrix[0])

	for i := 0; i < len(m.matrix); i++ {
		for j := 0; j < len(m.matrix[i]); j++ {
			m.matrix[i][j].neighbors = 0
		}
	}

	// loop over matrix
	for i := 0; i < len(m.matrix); i++ {
		for j := 0; j < len(m.matrix[i]); j++ {
			for _, neighbor := range neighbors {
				nx_cord = i + neighbor[0]
				ny_cord = j + neighbor[1]

				// check if neigbour coordinates are inside bounds
				if nx_cord >= 0 && ny_cord >= 0 && nx_cord < nrows && ny_cord < ncols {
					if m.matrix[nx_cord][ny_cord].curr_gen {
						m.matrix[i][j].neighbors += 1
					}
				}
			}
		}
	}
}

func (m model) changeGen() {
	nrows := len(m.matrix)
	ncols := len(m.matrix[0])

	for i := range nrows {
		for j := range ncols {
			curr_gen := &m.matrix[i][j].curr_gen
			next_gen := &m.matrix[i][j].next_gen
			neighbors := &m.matrix[i][j].neighbors

			if *curr_gen {
				// Currently alive cell
				if *neighbors < 2 {
					// Death by isolation: Fewer than 2 neighbors
					*next_gen = false
				} else if *neighbors == 2 || *neighbors == 3 {
					// Survival: 2 or 3 neighbors
					*next_gen = true
				} else {
					// Death by overcrowding: More than 3 neighbors
					*next_gen = false
				}
			} else {
				// Currently dead cell
				if *neighbors == 3 {
					// Birth: Exactly 3 neighbors
					*next_gen = true
				} else {
					// Remain dead
					*next_gen = false
				}
			}
			*next_gen, *curr_gen = *curr_gen, *next_gen
		}
	}
}
