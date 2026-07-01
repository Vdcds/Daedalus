// Package page provides a reusable page layout.
package page

import (
	"github.com/charmbracelet/lipgloss"
)

type Page struct {
	Header string
	Body   string
	Footer string
}

func New(header, body, footer string) *Page {
	return &Page{
		Header: header,
		Body:   body,
		Footer: footer,
	}
}

func (p *Page) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,

		p.Header,

		"",

		p.Body,

		"",

		p.Footer,
	)
}
