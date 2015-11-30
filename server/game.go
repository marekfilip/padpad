package server

import (
	"golang.org/x/net/websocket"
	"time"
)

type Game struct {
	Player1 *Client
	Player2 *Client
}

func NewGame() *Game {
	return &Game{nil, nil}
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
	var x, y uint = 0, 0
	for {
		x += 1
		y += 1

		ball := struct {
			x uint
			y uint
		}{
			x, y,
		}
		if g.Player1 != nil {
			websocket.JSON.Send(g.Player1.WebService, &ball)
		}
		if g.Player2 != nil {
			websocket.JSON.Send(g.Player2.WebService, &ball)
		}

		time.Sleep(time.Duration(time.Second / 60))
	}
}

type Games struct {
	content map[int]*Game
	nextId  int
}

func (games *Games) AddGame(g *Game) {
	games.content[games.nextId] = g
	games.nextId += 1
}

func (g *Games) Len() int {
	return len((*g).content)
}

func (g *Games) GetNextId() int {
	g.nextId += 1
	return (g.nextId - 1)
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
