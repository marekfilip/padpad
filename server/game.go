package server

import (
	"padpad/server/objects"
	"time"
)

type Game struct {
	Player1 *Client
	Player2 *Client
	Height  int
	Width   int
}

func NewGame() *Game {
	return &Game{nil, nil, 400, 300}
}

func (g *Game) AddPlayer(c *Client) bool {
	if g.Player1 == nil {
		g.Player1 = c
		return true
	}
	if g.Player2 == nil {
		g.Player2 = c
		return true
	}

	return false
}

func (g *Game) Start() {
	var b *objects.Ball = objects.NewBall(0, 0, 400, 300)
	for {
		b.Update()
		if g.Player1 != nil {
			g.Player1.ch <- b
		}
		if g.Player2 != nil {
			g.Player2.ch <- b
		}

		if g.Player1 == nil && g.Player1 == nil {
			// Zniszcz tą grę
		}
		time.Sleep(time.Duration(time.Second / 60))
	}
}

type Games struct {
	content   map[int]*Game
	currentId int
}

func (games *Games) AddGame(g *Game) {
	games.content[games.currentId] = g
	games.currentId += 1
}

func (g *Games) Len() int {
	return len((*g).content)
}

func (g *Games) AssignId() int {
	g.currentId++
	return g.currentId - 1
}

func (g *Games) Shift() *Game {
	if len((*g).content) == 0 {
		return nil
	}
	var x *Game
	for key, _ := range g.content {
		x = (*g).content[key]
		(*g).content[key] = nil
	}

	return x
}
