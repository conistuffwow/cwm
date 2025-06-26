package wm

import (
	"github.com/gdamore/tcell/v2"
)

type Window struct {
	X, Y          int
	Width, Height int
	Title         string
	Focused       bool
}

func NewWindow(x, y, w, h int, title string) *Window {
	return &Window{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
		Title:  title,
	}
}

func (w *Window) Draw(s tcell.Screen) { // this shit was ANNOYING to code.
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	if w.Focused {
		style = style.Background(tcell.ColorBlue)
	}
	// Draw corners
	s.SetContent(w.X, w.Y, '┌', nil, style)
	s.SetContent(w.X+w.Width-1, w.Y, '┐', nil, style)
	s.SetContent(w.X, w.Y+w.Height-1, '└', nil, style)
	s.SetContent(w.X+w.Width-1, w.Y+w.Height-1, '┘', nil, style)

	// Draw borders

	for cx := w.X + 1; cx < w.X+w.Width-1; cx++ {
		s.SetContent(cx, w.Y, '─', nil, style)
		s.SetContent(cx, w.Y+w.Height-1, '─', nil, style)
	}
	for cy := w.Y + 1; cy < w.Y+w.Height-1; cy++ {
		s.SetContent(w.X, cy, '│', nil, style)
		s.SetContent(w.X+w.Width-1, cy, '│', nil, style)
	}
	// Draw title bar
	for i, ch := range w.Title {
		if w.X+2+i < w.X+w.Width-1 {
			s.SetContent(w.X+2+i, w.Y, ch, nil, style)
		}
	}
	// all of that ^^ requires a STUPID amount of math.
}
