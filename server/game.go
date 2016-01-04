package server

import (
	"padpad/server/message"
	"padpad/server/objects"
	"time"
)

type Game struct {
	Id      int
	Queue   *Games
	Player1 *Client
	Player2 *Client
	Height  int
	Width   int
}

func NewGame(q *Games) *Game {
	return &Game{0, q, nil, nil, 400, 300}
}

func (g *Game) AddPlayer(c *Client) bool {
	if g.Player1 == nil {
		g.Player1 = c
		c.Pad.UpdatePadPos(float32(g.Player1.CanvasWidth/2-g.Player1.Pad.Length/2), float32(15))
		return true
	}
	if g.Player2 == nil {
		g.Player2 = c
		c.Pad.UpdatePadPos(float32(g.Player2.CanvasWidth/2-g.Player2.Pad.Length/2), float32(g.Player2.CanvasHeight-15))
		return true
	}

	return false
}

func (g *Game) Start() {
	var b *objects.Ball = objects.NewBall(0, 0, g.Height, g.Width)
	var forP1 [3]*message.Message
	var forP2 [3]*message.Message

	for {
		b.Update()
		if g.Player1 != nil {
			forP1[0] = b.Encode()
			forP1[1] = g.Player1.Pad.Encode(message.PLAYER_PAD_POSITION_TYPE)
			//g.Player1.ch <- b
			if g.Player2 != nil {
				//g.Player1.opponentPad <- g.Player2.Pad
				forP1[2] = g.Player2.Pad.Encode(message.OPPONENT_PAD_POSITION_TYPE)
			}
			g.Player1.toSend <- forP1
		}
		if g.Player2 != nil {
			forP2[0] = b.Encode()
			forP2[1] = g.Player2.Pad.Encode(message.PLAYER_PAD_POSITION_TYPE)
			//g.Player2.ch <- b
			if g.Player1 != nil {
				//g.Player2.opponentPad <- g.Player1.Pad
				forP2[2] = g.Player1.Pad.Encode(message.OPPONENT_PAD_POSITION_TYPE)
			}
			g.Player2.toSend <- forP2
		}

		/*if g.Player1 == nil || g.Player2 == nil {
			return
		}*/
		time.Sleep(time.Duration(time.Second / 60))
	}
}
func (g *Game) BothPlayersPresent() bool {
	if g.Player1 != nil && g.Player2 != nil {
		return true
	}
	return false
}

type Games struct {
	content   map[int]*Game
	currentId int
}

func (games *Games) AddGame(g *Game) {
	g.Id = games.currentId
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
		if x.Player1 == nil || x.Player2 == nil {
			return x
		}
	}

	return nil
}

func (g *Games) Remove(game *Game) {
	delete(g.content, game.Id)
}
