package main

import (
	b "game/ball"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
	Ball   b.Ball
}

func main() {

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	ball := b.Ball{}

	game := Game{
		screen: screen,
		Ball:   ball.Initiate(),
	}

	go game.Run()

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)

	g.screen.SetStyle(defStyle)

	x := 0
	for {
		// g.screen.Clear()
		width, height := g.screen.Size()
		g.Ball.CheckEdge(width, height)
		g.Ball.Update()

		g.screen.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		g.screen.Show()
		x++

		time.Sleep(5 * time.Millisecond)
	}
}
