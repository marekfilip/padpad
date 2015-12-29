package objects

import "padpad/server/message"
import "fmt"

type Pad struct {
	X      float32
	Y      float32
	Length float32
}

func NewPad(x, y float32) *Pad {
	return &Pad{x, y, 0}
}

func (p *Pad) UpdatePadPos(x, y float32) {
	fmt.Println("Updating pad position X:", p.X, "Y:", p.Y)
	p.X = x
	p.Y = y
}

func (p *Pad) UpdatePadLength(l float32) {
	p.Length = l
}

func (p *Pad) Encode(msgType int) *message.Message {
	return &message.Message{
		msgType,
		map[string]float32{
			"x": p.X,
			"y": p.Y,
		},
	}
}
