package server

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	pattern        string
	WaitingClients *Queue
	WaitingGames   []*Game
	Games          []*Game
	addCh          chan *Client
	delCh          chan *Client
	doneCh         chan bool
	errCh          chan error
}

func NewServer(pattern string) *Server {
	waitingClients := new(Queue)
	waitingGames := []*Game{}
	Games := []*Game{}
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		waitingClients,
		waitingGames,
		Games,
		addCh,
		delCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) Listen() {
	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.WaitingClients.Add(c)
			log.Println("Now", len(*s.WaitingClients), "clients waiting.")

			// del a client
			/*case c := <-s.delCh:
				log.Println("Delete client")
				delete(s.clients, c.id)

			case err := <-s.errCh:
				log.Println("Error:", err.Error())

			case <-s.doneCh:
				return
			*/
		}
	}
}
