package wm

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Manager struct {
	Windows []*Window
	Focus   int // index of shit idk

	screen tcell.Screen
}

func NewManager() *Manager {
	return &Manager{
		Windows: []*Window{},
		Focus:   -1,
	}
}

func (m *Manager) AddWindow(w *Window, sW, sH int, s tcell.Screen) {
	m.Windows = append(m.Windows, w)
	if m.Focus == -1 {
		m.Focus = 0
		m.Windows[0].Focused = true
	}
	m.LayoutTiles(sW, sH)
	m.DrawPanel()
}

func (m *Manager) Draw(s tcell.Screen) {
	for i, w := range m.Windows {
		w.Focused = (i == m.Focus)
		w.Draw(s)
		m.DrawPanel()

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
		w.Y = 1
		w.Width = tileWidth
		w.Height = sH - 1
	}
}
func (m *Manager) DrawPanel() {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue)
	width, _ := m.screen.Size()
	s := m.screen
	for x := 0; x < width; x++ {
		s.SetContent(x, 0, ' ', nil, style)
	}

	// left
	title := "No Window"
	if m.Focus >= 0 && m.Focus < len(m.Windows) {
		title = m.Windows[m.Focus].Title
	}
	drawText(s, 1, 0, style, title)

	// right
	timeStr := time.Now().Format("15:04:05")
	battery := getBatteryStatus()
	statusTxt := fmt.Sprintf("%s  %s", battery, timeStr)
	xStart := width - len(statusTxt) - 1
	drawText(s, xStart, 0, style, statusTxt)
}

func drawText(s tcell.Screen, x, y int, style tcell.Style, text string) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}
func getBatteryStatus() string {
	return "placeholder battery"
}
func (m *Manager) SetScreen(s tcell.Screen) {
	m.screen = s
}
