// Package theme defines Daedalus' theming system.
package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name    string
	Palette Palette
	Styles  Styles
}

type Palette struct {
	Base    lipgloss.Color
	Surface lipgloss.Color
	Overlay lipgloss.Color

	Text  lipgloss.Color
	Muted lipgloss.Color

	Love lipgloss.Color
	Gold lipgloss.Color
	Rose lipgloss.Color
	Pine lipgloss.Color
	Foam lipgloss.Color
	Iris lipgloss.Color
}
