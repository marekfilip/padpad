package server

type Queue struct {
	content   map[int]*Client
	currentId int
}

func NewQueue() *Queue {
	return &Queue{
		make(map[int]*Client),
		0,
	}
}

func (q *Queue) Add(c *Client) {
	q.content[c.Id] = c
}

func (q *Queue) Len() int {
	return len((*q).content)
}

func (q *Queue) AssignId() int {
	q.currentId++
	return q.currentId - 1
}

func (q *Queue) Remove(c *Client) {
	delete(q.content, c.Id)
}
