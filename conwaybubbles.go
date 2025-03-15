package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tile struct {
	curr_gen  bool
	next_gen  bool
	neighbors int
}

type model struct {
	matrix    [][]tile
	cursor    map[string]int
	altscreen bool
}

func (m model) toggleState(x int, y int) {
	m.matrix[x][y].curr_gen = !m.matrix[x][y].curr_gen
}

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

func initialModel() model {
	matrix := make([][]tile, 10)
	for row := range matrix {
		matrix[row] = make([]tile, 10)
	}

	return model{
		matrix: matrix,
		cursor: map[string]int{"x": 0, "y": 0},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")) // Red text
		// Background(lipgloss.Color("#FFFF00")). // Yellow background
		// Padding(0, 1)

	cursorStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7d33ff"))

	deadStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#000000"))

	dieStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FF0000"))

	birthStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90ee90"))

	surviveStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90eeee"))
)

func (m model) RenderHeader() string {

	header := "Conway's Game Of Life\n"
	header += fmt.Sprintf("Cursor pos: x:%d y:%d\n", m.cursor["x"], m.cursor["y"])
	header += fmt.Sprintf("Map Size: x:%d y:%d\n", len(m.matrix), len(m.matrix[0]))
	header += "(c) count neighbors, (a) toggle between map and neighbor count view"
	return headerStyle.Render(header)
}

func (m model) RenderGameMap() string {
	var game_map string
	for i := range m.matrix {
		game_map += "\n"
		for j := range m.matrix[i] {
			var cell string
			if m.matrix[i][j].curr_gen {
				cell = "👽"
				// cell = "o"
			} else {
				cell = "💀"
				// cell = "x"
			}

			// cursor highlighter
			if i == m.cursor["y"] && j == m.cursor["x"] {
				cell = cursorStyle.Render(cell)
			}
			game_map += cell
		}
	}
	return game_map
}

func (m model) RenderNeighborCount() string {
	var neighbor_count string
	var styled_cell string
	for i := range m.matrix {
		neighbor_count += "\n"
		for j := range m.matrix[i] {

			cell := m.matrix[i][j]
			cellText := fmt.Sprintf(" %d", cell.neighbors)

			if cell.curr_gen {
				// Currently alive cell
				if cell.neighbors < 2 {
					// Death by isolation: Fewer than 2 neighbors
					styled_cell = dieStyle.Render(cellText)
				} else if cell.neighbors == 2 || cell.neighbors == 3 {
					// Survival: 2 or 3 neighbors
					styled_cell = surviveStyle.Render(cellText)
				} else {
					// Death by overcrowding: More than 3 neighbors
					styled_cell = dieStyle.Render(cellText)
				}
			} else {
				// Currently dead cell
				if cell.neighbors == 3 {
					// Birth: Exactly 3 neighbors
					styled_cell = birthStyle.Render(cellText)
				} else {
					// Remain dead
					styled_cell = deadStyle.Render(cellText)
				}
			}

			// render cursor
			if i == m.cursor["y"] && j == m.cursor["x"] {
				styled_cell = cursorStyle.Render(cellText)
			}
			neighbor_count += styled_cell
		}
	}
	return neighbor_count
}

func (m model) View() string {
	header := m.RenderHeader()

	var content string
	if !m.altscreen {
		content = m.RenderGameMap()
	} else {
		content = m.RenderNeighborCount()
	}
	return header + content + "\n"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// quit
		case "q", "ctr+c":
			return m, tea.Quit
		// move cursor
		case "up":
			if m.cursor["y"] > 0 {
				m.cursor["y"] += -1
			}
		case "down":
			if m.cursor["y"] < len(m.matrix[0])-1 {
				m.cursor["y"] += 1
			}
		case "left":
			if m.cursor["x"] > 0 {
				m.cursor["x"] += -1
			}
		case "right":
			if m.cursor["x"] < len(m.matrix)-1 {
				m.cursor["x"] += 1
			}
		// switch tile state
		case "e":
			m.toggleState(m.cursor["y"], m.cursor["x"])
		// next generation
		case "c":
			// execute count and change gen
			m.countNeighbors()
		case "a":
			m.altscreen = !m.altscreen
		}

	}
	return m, nil
}
