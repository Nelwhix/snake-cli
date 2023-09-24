package main

import (
	"github.com/Nelwhix/snake-cli/internal/game"
	"github.com/Nelwhix/snake-cli/internal/snake"
	"github.com/gdamore/tcell"
	"log"
	"os"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	snakeParts := []snake.Part{
		{
			X: 5,
			Y: 10,
		},
		{
			X: 6,
			Y: 10,
		},
		{
			X: 7,
			Y: 10,
		},
	}

	snakeBody := snake.NewBody(snakeParts, 1, 0)
	nGame := game.NewGame(screen, snakeBody)

	go nGame.Run()
	for {
		switch event := nGame.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			nGame.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				nGame.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				nGame.SnakeBody.ChangeDir(-1, 0)
			} else if event.Key() == tcell.KeyDown {
				nGame.SnakeBody.ChangeDir(1, 0)
			} else if event.Key() == tcell.KeyLeft {
				nGame.SnakeBody.ChangeDir(0, -1)
			} else if event.Key() == tcell.KeyRight {
				nGame.SnakeBody.ChangeDir(0, 1)
			}
		}
	}
}
