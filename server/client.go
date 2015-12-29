package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"padpad/server/message"
	"padpad/server/objects"
)

const channelBufSize = 100

type Client struct {
	Id           int
	Points       uint
	CanvasHeight float32
	CanvasWidth  float32
	WebService   *websocket.Conn
	Server       *Server
	Game         *Game
	Pad          *objects.Pad
	doneCh       chan bool
	ch           chan *objects.Ball
	playerPad    chan *objects.Pad
	opponentPad  chan *objects.Pad
}

func NewClient(id int, ws *websocket.Conn, server *Server) *Client {
	return &Client{
		id,
		0,
		0,
		0,
		ws,
		server,
		nil,
		objects.NewPad(0, 0),
		make(chan bool),
		make(chan *objects.Ball, channelBufSize),
		make(chan *objects.Pad, channelBufSize),
		make(chan *objects.Pad, channelBufSize),
	}
}

func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) StartGame() bool {
	c.Server.StartGame(c)
	return true
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {
		// send opponent pad position
		case msg := <-c.opponentPad:
			websocket.JSON.Send(c.WebService, msg.Encode(message.OPPONENT_PAD_POSITION_TYPE))
		// clients pad position
		case msg := <-c.playerPad:
			websocket.JSON.Send(c.WebService, msg.Encode(message.PLAYER_PAD_POSITION_TYPE))
		// send ball position to client
		case msg := <-c.ch:
			websocket.JSON.Send(c.WebService, msg.Encode())
		// receive done request
		case <-c.doneCh:
			if c.Game != nil {
				c.Game.Player1 = nil
				c.Game.Player2 = nil
				c.Server.delGameChan <- c.Game
			}
			c.Server.DelClient(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")

	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.Server.DelClient(c)
			c.doneCh <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg message.Message
			err := websocket.JSON.Receive(c.WebService, &msg)
			c.Decode(&msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.Server.Err(err)
			}
		}
	}
}

/*
	Dokończyć metodę dekodującą dane przychodzące
*/

func (c *Client) Decode(msg *message.Message) {
	switch {
	case msg.MessageType == message.STARTGAME_TYPE:
		fmt.Println(c.Id, "received startgame:", *msg)
		/*
			Rozmiary planszy gracza nie mają w tej chwili znaczenia
		*/
		c.CanvasHeight = msg.Data["cH"]
		c.CanvasWidth = msg.Data["cW"]
		c.Pad.UpdatePadLength(c.CanvasWidth / 8)
		c.StartGame()
		fmt.Println("Player:", c.Pad)
		websocket.JSON.Send(c.WebService, c.Pad.Encode(message.PLAYER_PAD_POSITION_TYPE))
		break
	case msg.MessageType == message.PAD_POSITION_TYPE:
		var x float32 = msg.Data["pX"]
		//fmt.Println(c.Id, "received podposition:", *msg)

		if x < 0 {
			x = 0
		}

		if (x + c.Pad.Length) > c.CanvasWidth {
			x = c.CanvasWidth - c.Pad.Length
		}

		c.Pad.UpdatePadPos(x, c.Pad.Y)
		c.playerPad <- c.Pad
		break
	default:
		fmt.Println("Unrecognized msg:", *msg)
	}
}
