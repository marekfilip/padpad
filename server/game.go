package server

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
