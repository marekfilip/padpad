package server

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	pattern       string
	Clients       *Queue
	Games         *Games
	addClientChan chan *Client
	delClientChan chan *Client
	startCh       chan *Client
	/*addGameChan   chan *Game*/
	delGameChan chan *Game
	doneCh      chan bool
	errCh       chan error
}

func NewServer(pattern string) *Server {
	clients := &Queue{make(map[int]*Client), 0}
	Games := &Games{make(map[int]*Game), 0}
	addClientChan := make(chan *Client)
	delClientChan := make(chan *Client)
	startCh := make(chan *Client)
	delGameChan := make(chan *Game)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		clients,
		Games,
		addClientChan,
		delClientChan,
		startCh,
		delGameChan,
		doneCh,
		errCh,
	}
}

func (s *Server) AddClient(c *Client) {
	s.addClientChan <- c
}

func (s *Server) DelClient(c *Client) {
	s.delClientChan <- c
}

func (s *Server) StartGame(c *Client) {
	s.startCh <- c
}

func (s *Server) DelGame(g *Game) {
	s.delGameChan <- g
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

		client := NewClient(s.Clients.AssignId(), ws, s)
		s.AddClient(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	var tempGame *Game
	for {
		select {
		// Add new a client
		case c := <-s.addClientChan:
			s.Clients.Add(c)
			log.Println("Added new client. Now", s.Clients.Len(), "clients online.")
		// del a client
		case c := <-s.delClientChan:
			s.Clients.Remove(c)
			log.Println("Delete client. Now", s.Clients.Len(), "clients online.")
		case g := <-s.delGameChan:
			s.Games.Remove(g)
			log.Println("Delete game. Now", s.Games.Len(), "games online.")
		case c := <-s.startCh:
			tempGame = s.Games.Shift()
			if tempGame == nil {
				tempGame = NewGame(s.Games)
				s.Games.AddGame(tempGame)
			}
			c.Game = tempGame
			c.Game.AddPlayer(c)
			if c.Game.BothPlayersPresent() {
				go c.Game.Start()
			}
			log.Println("Starting new game. Now", s.Games.Len(), "games online.")
		case err := <-s.errCh:
			log.Println("Error:", err.Error())
		case <-s.doneCh:
			return
		}
	}
}
