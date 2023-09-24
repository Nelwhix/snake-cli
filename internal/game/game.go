package game

import (
	"github.com/Nelwhix/snake-cli/internal/snake"
	"github.com/gdamore/tcell"
	"time"
)

type Game struct {
	Screen    tcell.Screen
	SnakeBody *snake.Body
}

func NewGame(screen tcell.Screen, body *snake.Body) *Game {
	return &Game{
		Screen:    screen,
		SnakeBody: body,
	}
}

func drawParts(s tcell.Screen, parts []snake.Part, style tcell.Style) {
	for _, part := range parts {
		s.SetContent(part.X, part.Y, ' ', nil, style)
	}
}

func (g *Game) Run() {
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		g.Screen.Clear()
		g.SnakeBody.Update(width, height)
		drawParts(g.Screen, g.SnakeBody.Parts, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}
}
