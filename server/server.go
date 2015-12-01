package server

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	"padpad/objects"
)

type Server struct {
	pattern        string
	WaitingClients *Queue
	WaitingGames   *Games
	Games          *Games
	addCh          chan *Client
	startCh        chan *Client
	delCh          chan *Client
	doneCh         chan bool
	errCh          chan error
}

func NewServer(pattern string) *Server {
	waitingClients := &Queue{make(map[int]*Client), 0}
	waitingGames := &Games{make(map[int]*Game), 0}
	Games := &Games{make(map[int]*Game), 0}
	addCh := make(chan *Client)
	startCh := make(chan *Client)
	delCh := make(chan *Client)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &Server{
		pattern,
		waitingClients,
		waitingGames,
		Games,
		addCh,
		startCh,
		delCh,
		doneCh,
		errCh,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) StartGame(c *Client) {
	s.startCh <- c
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

		client := NewClient(s.WaitingClients.GetNextId(), ws, s)
		s.Add(client)
		client.Listen()
		client.ch <- &objects.Ball{1, 1}
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler")

	var tempGame *Game
	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.WaitingClients.Add(c)
			log.Println("Now", s.WaitingClients.Len(), "clients waiting.")

		case c := <-s.addCh:
			log.Println("Starting new game")
			if s.WaitingGames.Len() > 0 {
				tempGame = s.WaitingGames.Shift()
				tempGame.AddPlayer(c)
				s.Games.AddGame(tempGame)
			} else {
				tempGame = NewGame()
				tempGame.AddPlayer(c)
				go tempGame.Start()
				s.WaitingGames.AddGame(tempGame)
			}
		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.WaitingClients.content, c.Id)
			/*
				case err := <-s.errCh:
					log.Println("Error:", err.Error())

				case <-s.doneCh:
					return
			*/
		}
	}
}
