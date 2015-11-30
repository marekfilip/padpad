package server

type Queue []*Client

func (q *Queue) Add(c *Client) {
	*q = append(*q, c)
}
