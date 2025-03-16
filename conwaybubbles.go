package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
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
	textInput textinput.Model
	altscreen bool
	nrows     int
	ncols     int
}

// func (m *model) initMatrix() {
// 	m.matrix = make([][]tile, 10)
// 	for row := range m.matrix {
// 		m.matrix[row] = make([]tile, 10)
// 	}
// }

func initialModel(nrows int, ncols int) model {
	matrix := make([][]tile, nrows)
	for row := range matrix {
		matrix[row] = make([]tile, ncols)
	}

	return model{
		matrix: matrix,
		cursor: map[string]int{"x": 0, "y": 0},
		nrows:  nrows,
		ncols:  ncols,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel(10, 50))
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
	header += "(e) toggle tile state, (n) activate next generation\n"
	header += "(a) toggle between map and neighbor count view, (q) quit"
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
			if i == m.cursor["x"] && j == m.cursor["y"] {
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

			cell := &m.matrix[i][j]
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
			if i == m.cursor["x"] && j == m.cursor["y"] {
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
			if m.cursor["x"] > 0 {
				m.cursor["x"] += -1
			}
		case "down":
			if m.cursor["x"] < m.nrows-1 {
				m.cursor["x"] += 1
			}
		case "left":
			if m.cursor["y"] > 0 {
				m.cursor["y"] += -1
			}
		case "right":
			if m.cursor["y"] < m.ncols-1 {
				m.cursor["y"] += 1
			}
		// switch tile state
		case "e":
			m.toggleState(m.cursor["x"], m.cursor["y"])
		// next generation
		case "c":
			// execute count and change gen
			m.countNeighbors()
		case "a":
			m.altscreen = !m.altscreen
			// go to next gen
		case "n":
			m.countNeighbors()
			m.changeGen()
			m.RenderGameMap()
		}

	}
	return m, nil
}
