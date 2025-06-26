package main

import (
	"log"
	"time"

	"github.com/conistuffwow/cwm/util"
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
	screen.EnableMouse()
	defer screen.Fini()

	screen.Clear()

	manager := wm.NewManager()
	manager.SetScreen(screen)
	util.AddWindowWithLayout(manager, screen, "Wnd1")
	util.AddWindowWithLayout(manager, screen, "Wnd2")
	util.AddWindowWithLayout(manager, screen, "Wnd3")
	manager.Draw(screen)
	screen.Show()
	eventCh := make(chan tcell.Event, 10)
	go func() {
		for {
			ev := screen.PollEvent()
			eventCh <- ev
		}
	}()
	drawTicker := time.NewTicker(time.Second)
	defer drawTicker.Stop()
	for {
		select {
		case <-drawTicker.C:
			manager.DrawPanel()
			screen.Show()
		case ev := <-eventCh:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyCtrlC:
					return
				}
				manager.HandleEvent(ev, screen)
				screen.Clear()
				manager.Draw(screen)
				manager.DrawPanel()
				screen.Show()
			}
		}
	}

}
