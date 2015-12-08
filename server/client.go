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
	CanvasHeight int
	CanvasWidth  int
	WebService   *websocket.Conn
	Server       *Server
	Game         *Game
	Pad          *objects.Pad
	doneCh       chan bool
	ch           chan *objects.Ball
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

func (c *Client) StartGame() {
	c.Server.StartGame(c)
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			//log.Println("Send:", *msg)
			websocket.JSON.Send(c.WebService, msg.Encode())
		// receive done request
		case <-c.doneCh:
			c.Server.Del(c)
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
			c.Server.Del(c)
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
		fmt.Println("Recived startgame:", *msg)
		c.CanvasHeight = int(msg.Data["cH"])
		c.CanvasHeight = int(msg.Data["cW"])
		c.Pad.UpdatePadLength(msg.Data["cW"] / 8)
		c.StartGame()
		break
	case msg.MessageType == message.PAD_POSITION_TYPE:
		fmt.Println("Recived podposition:", *msg)
		c.Pad.UpdatePadPos(msg.Data["pX"])
		break
	default:
		fmt.Println("Unrecognized msg:", *msg)
	}
}
