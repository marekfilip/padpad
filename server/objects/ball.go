package objects

import "padpad/server/message"

type Ball struct {
	X uint
	Y uint
}

func NewBall(startx, starty uint) *Ball {
	return &Ball{
		startx,
		starty,
	}
}

func (b *Ball) Update() {
	b.X++
	b.Y++
}

func (b *Ball) Encode() *message.Message {
	return &message.Message{
		message.BALL_POSITION_TYPE,
		map[string]interface{}{
			"x": b.X,
			"y": b.Y,
		},
	}
}
