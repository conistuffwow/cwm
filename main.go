package main

import (
	"log"

	"github.com/conistuffwow/cwm/wm"
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}
	defer screen.Fini()

	screen.Clear()

	manager := wm.NewManager()
	manager.AddWindow(wm.NewWindow(5, 3, 20, 10, "First Window"))
	manager.AddWindow(wm.NewWindow(30, 5, 25, 12, "Second Window"))

	manager.Draw(screen)
	screen.Show()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				return
			case tcell.KeyTab:
				manager.FocusNext()
				screen.Clear()
				manager.Draw(screen)
				screen.Show()
			}
		}
	}

}
