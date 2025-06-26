package util

import (
	"github.com/conistuffwow/cwm/wm"
	"github.com/gdamore/tcell/v2"
)

func AddWindowWithLayout(m *wm.Manager, s tcell.Screen, title string) {
	w := wm.NewWindow(0, 0, 0, 0, title)
	width, height := s.Size()
	m.AddWindow(w, width, height, s)

}
