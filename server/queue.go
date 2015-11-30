package server

type Queue struct {
	content map[int]*Client
	nextId  int
}

func NewQueue() *Queue {
	return &Queue{
		make(map[int]*Client),
		0,
	}
}

func (q *Queue) Add(c *Client) {
	q.content[q.nextId] = c
	q.nextId += 1
}

func (q *Queue) Len() int {
	return len((*q).content)
}

func (q *Queue) GetNextId() int {
	q.nextId += 1
	return (q.nextId - 1)
}
