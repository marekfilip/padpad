package server

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	Id         int
	Points     uint
	WebService *websocket.Conn
	Server     *Server
}

func NewClient(id int, ws *websocket.Conn, server *Server) *Client {
	return &Client{id, 0, ws, server}
}
