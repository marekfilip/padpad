package server

type Client struct {
	Points uint
}

func NewClient() *Client {
	return &Client{}
}
