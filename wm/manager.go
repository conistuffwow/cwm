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

func (m *Manager) AddWindow(w *Window, sW, sH int) {
	m.Windows = append(m.Windows, w)
	if m.Focus == -1 {
		m.Focus = 0
		m.Windows[0].Focused = true
	}
	m.LayoutTiles(sW, sH)
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

func (m *Manager) FocusWindow(index int) {
	if index >= 0 && index < len(m.Windows) {
		if m.Focus != -1 {
			m.Windows[m.Focus].Focused = true
		}
		m.Focus = index
		m.Windows[m.Focus].Focused = true
	}
}

func (m *Manager) HandleEvent(ev tcell.Event, s tcell.Screen) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyTab, tcell.KeyRight:
			m.FocusNext()
		case tcell.KeyLeft:
			m.FocusPrev()
		}

	}
}

func (m *Manager) FocusPrev() {
	if len(m.Windows) == 0 {
		return
	}
	m.Windows[m.Focus].Focused = false
	m.Focus = (m.Focus - 1 + len(m.Windows)) % len(m.Windows)
	m.Windows[m.Focus].Focused = true
}

func (m *Manager) LayoutTiles(sW, sH int) {
	count := len(m.Windows)
	if count == 0 {
		return
	}
	tileWidth := sW / count
	for i, w := range m.Windows {
		w.X = i * tileWidth
		w.Y = 0
		w.Width = tileWidth
		w.Height = sH
	}
}
