package server

import (
	"golang.org/x/net/websocket"
)

type Client struct {
	Points     uint
	WebService *websocket.Conn
	Server     *Server
}

func NewClient(ws *websocket.Conn, server *Server) *Client {
	return &Client{0, ws, server}
}
