package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tile struct {
	curr_gen  bool
	next_gen  bool
	neighbors int
}

type model struct {
	// layout     layout

	// state
	matrix     [][]tile
	cursor     map[string]int
	altscreen  bool
	autorun    bool
	nrows      int
	ncols      int
	generation int
	mousex     int
	mousey     int
}

func initialModel(nrows int, ncols int) model {
	matrix := make([][]tile, nrows)
	for row := range matrix {
		matrix[row] = make([]tile, ncols)
	}

	return model{
		matrix:     matrix,
		cursor:     map[string]int{"x": 0, "y": 0},
		nrows:      nrows,
		ncols:      ncols,
		generation: 0,
		autorun:    false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel(15, 50))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch/ distinguish between message types like Tick or KeyMsg
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
			m.countNeighbors()
		case "r":
			// reset counter
			m.generation = 0
			// case "x":
			// 	// move forward x generations
			// 	m.RenderModal()
		case "s":
			// start stop autorun
			m.autorun = !m.autorun

			// if autorun start autorunCmd
			if m.autorun {
				return m, autorunCmd()
			}
		}
	case AutorunMsg:
		if m.autorun {
			m.countNeighbors()
			m.changeGen()
			m.RenderGameMap()
			m.countNeighbors()

			return m, autorunCmd()
		}
	}
	// m.textInput, cmd = m.textInput.Update(msg)
	return m, nil
}
func autorunCmd() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return AutorunMsg(t)
		// return {t}
	})
}

type AutorunMsg time.Time

// func (m *model) autorun_func() {
// 	for m.autorun {
// 		time.Sleep(500 * time.Millisecond)
// 		m.countNeighbors()
// 		m.changeGen()
// 		m.RenderGameMap()
// 		m.countNeighbors()
// 		if !m.autorun {
// 			break
// 		}
// 	}
// }
