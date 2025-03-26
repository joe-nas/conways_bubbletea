package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Green    = lipgloss.Color("#b6f486")
	Blue     = lipgloss.Color("#88d4c3")
	Purple1  = lipgloss.Color("#a8aedd")
	Purlple2 = lipgloss.Color("#8769b6")
	Purple3  = lipgloss.Color("#400e63")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Underline(true).
			Height(2).
			Width(30)

	HeaderStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#400e63")).
			Width(10).
			Height(10).BorderStyle(lipgloss.RoundedBorder())

	HelpStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#88d4c3")).
			Width(10).
			Height(10).
			BorderStyle(lipgloss.RoundedBorder())

	BodyStyle = lipgloss.NewStyle().
		// Background(lipgloss.Color("#b6f486")).
		Width(120).
		Height(20).
		BorderStyle(lipgloss.RoundedBorder())

	FooterStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#a8aedd")).
			Width(10).
			Height(7).BorderStyle(lipgloss.RoundedBorder())
)

var (
	// Overall cursor style
	CursorStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#7d33ff"))
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
