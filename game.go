package main

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
	Ball   Ball
}

type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (b *Ball) Display() rune {
	return '\u25CF'
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) CheckEdge(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.Xspeed *= -1
	}
	if b.Y <= 0 || b.Y >= maxHeight {
		b.Yspeed *= -1
	}

}

func main() {

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	ball := Ball{
		X:      1,
		Y:      1,
		Xspeed: 1,
		Yspeed: 1,
	}
	game := Game{
		screen: screen,
		Ball:   ball,
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

		time.Sleep(40 * time.Millisecond)
	}
}
