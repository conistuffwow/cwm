package wm

import (
	"github.com/gdamore/tcell/v2"
)

type Manager struct {
	Windows []*Window
	Focus   int // index of shit idk
}

func NewManager() *Manager {
	return &Manager{
		Windows: []*Window{},
		Focus:   -1,
	}
}

func (m *Manager) AddWindow(w *Window) {
	m.Windows = append(m.Windows, w)
	if m.Focus == -1 {
		m.Focus = 0
		m.Windows[0].Focused = true
	}
}

func (m *Manager) Draw(s tcell.Screen) {
	for i, w := range m.Windows {
		w.Focused = (i == m.Focus)
		w.Draw(s)
	}
}

func (m *Manager) FocusNext() {
	if len(m.Windows) == 0 {
		return
	}
	m.Windows[m.Focus].Focused = false
	m.Focus = (m.Focus + 1) % len(m.Windows)
	m.Windows[m.Focus].Focused = true
}
