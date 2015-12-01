package server

import (
	"golang.org/x/net/websocket"
	//"io"
	"log"
	"padpad/objects"
)

const channelBufSize = 100

type Client struct {
	Id         int
	Points     uint
	WebService *websocket.Conn
	Server     *Server
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
	//c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			log.Println("Send:", *msg)
			websocket.JSON.Send(c.WebService, *msg)

		// receive done request
		case <-c.doneCh:
			c.Server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
/*func (c *Client) listenRead() {
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
			var msg objects.Ball
			err := websocket.JSON.Receive(c.WebService, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.Server.Err(err)
			} else {
				c.Server.SendAll(&msg)
			}
		}
	}
}
*/
