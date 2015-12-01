package server

import (
	//"fmt"
	//"golang.org/x/net/websocket"
	"padpad/objects"
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
	var b *objects.Ball = objects.NewBall(100, 100)
	for {
		b.X += 1
		b.Y += 1

		if g.Player1 != nil {
			//fmt.Println("Wysyłam do gracza 1", *b)
			g.Player1.ch <- b
			//websocket.JSON.Send(g.Player1.WebService, *b)
		}
		if g.Player2 != nil {
			//fmt.Println("Wysyłam do gracza 2", *b)
			g.Player2.ch <- b
			//websocket.JSON.Send(g.Player2.WebService, *b)
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
