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
	Points       uint16
	CanvasHeight float32
	CanvasWidth  float32
	WebService   *websocket.Conn
	Server       *Server
	Game         *Game
	Pad          *objects.Pad
	doneCh       chan bool
	toSend       chan [3]*message.Message
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
		make(chan [3]*message.Message, channelBufSize),
	}
}

func (c *Client) AddPoint() {
	c.Points++
}

func (c *Client) RemovePoint() {
	if c.Points != 0 {
		c.Points--
	}
}

func (c *Client) RemovePoints() {
	c.Points = 0
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
		case msgs := <-c.toSend:
			websocket.JSON.Send(c.WebService, msgs)
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
		break
	case msg.MessageType == message.INC_PAD_POSITION_TYPE:
		var x float32 = msg.Data["pX"]

		if x < 0 {
			x = 0
		}

		if x > c.CanvasWidth {
			x = c.CanvasWidth
		}

		c.Pad.UpdatePadPos(x, c.Pad.Y)
		break
	default:
		fmt.Println("Unrecognized msg:", *msg)
	}
}

func (c *Client) Encode(msgType int, canvasWidth float32, usePadLengthTransition bool) *message.Message {
	var x = c.Pad.X
	if usePadLengthTransition {
		x = c.Pad.X - (c.Pad.Length / 2)
		if x < 0 {
			x = 0
		}
		if x > (canvasWidth - c.Pad.Length) {
			x = canvasWidth - c.Pad.Length
		}
	}
	return &message.Message{
		msgType,
		map[string]float32{
			"x": x,
			"y": c.Pad.Y,
			"l": c.Pad.Length,
			"p": float32(c.Points),
		},
	}
}
