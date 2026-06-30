// Package rosepine provides the Rose Pine theme for Daedalus.
package rosepine

import "github.com/vdcds/Daedalus/internal/tui/theme"

func New() theme.Theme {
	palette := theme.Palette{
		Base:    "#191724",
		Surface: "#1f1d2e",
		Overlay: "#26233a",

		Text:  "#e0def4",
		Muted: "#908caa",

		Love: "#eb6f92",
		Gold: "#f6c177",
		Rose: "#ebbcba",
		Pine: "#31748f",
		Foam: "#9ccfd8",
		Iris: "#c4a7e7",
	}

	return theme.Theme{
		Name:    "Rose Pine",
		Palette: palette,
		Styles:  theme.NewStyles(palette),
	}
}
