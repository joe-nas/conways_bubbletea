package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

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
			m.countDeadAlive()
		// next generation
		case "c":
			// execute count and change gen
			m.countNeighbors()
		case "a":
			// start stop autorun
			m.autorun = !m.autorun
			// if autorun start autorunCmd
			if m.autorun {
				return m, autorunCmd()
			}
		case "n":
			m.countNeighbors()
			m.changeGen()
			m.renderGameMap()
			m.countNeighbors()
		case "r":
			// reset counter-game: write a function to reset the game
			m.autorun = false
			m.generation = 0
		// case "x":
		// 	// move forward x generations
		// 	m.RenderModal()
		case "s":
			// go to next gen
			m.altscreen = !m.altscreen
		}

	case AutorunMsg:
		if m.autorun {
			m.countNeighbors()
			m.changeGen()
			m.renderGameMap()
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
