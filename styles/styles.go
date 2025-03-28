package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Green   = lipgloss.Color("#b6f486")
	Blue    = lipgloss.Color("#88d4c3")
	Purple1 = lipgloss.Color("#a8aedd")
	Purple2 = lipgloss.Color("#8769b6")
	Purple3 = lipgloss.Color("#400e63")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(Green).
			Underline(true).
			Height(1).
			Width(30)

	StatsStyle = lipgloss.NewStyle().
			Background(Purple1).
			Width(33).
			Height(2)
		// .BorderStyle(lipgloss.RoundedBorder())

	HelpStyle = lipgloss.NewStyle().
			Background(Purple2).
			Width(99).
			Height(2)
		// BorderStyle(lipgloss.RoundedBorder())

	BodyStyle = lipgloss.NewStyle().
		// Background(lipgloss.Color("#b6f486")).
		Width(101).
		Height(15).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
		// BorderStyle(lipgloss.RoundedBorder())

	FooterStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#a8aedd")).
			Width(10).
			Height(7).BorderStyle(lipgloss.RoundedBorder())
)

var (
	// Overall cursor style
	CursorStyle = lipgloss.NewStyle().
		Background(Purple2)
)

var (

	// Styles for count view
	DeadStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#000000"))

	DieStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#FF0000"))

	BirthStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90ee90"))

	SurviveStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#90eeee"))
)
