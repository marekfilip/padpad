package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"padpad/server/objects"
)

const channelBufSize = 100

type Client struct {
	Id         int
	Points     uint
	WebService *websocket.Conn
	Server     *Server
	Game       *Game
	doneCh     chan bool
	ch         chan *objects.Ball
}

func NewClient(id int, ws *websocket.Conn, server *Server) *Client {
	return &Client{id, 0, ws, server, make(chan bool), make(chan *objects.Ball, channelBufSize)}
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
			log.Println("Send:", *msg)
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
			var msg Message
			err := websocket.JSON.Receive(c.WebService, &msg)
			fmt.Println("Recived message:", msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.Server.Err(err)
			} /*else {
				c.Server.SendAll(&msg)
			}*/
			fmt.Println("Default sie odpalil")
		}
	}
}
