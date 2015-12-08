package objects

type Pad struct {
	X      float32
	Y      float32
	Length float32
}

func NewPad(x, y float32) *Pad {
	return &Pad{x, y, 0}
}

func (p *Pad) UpdatePadPos(x float32) {
	p.X = x
}

func (p *Pad) UpdatePadLength(l float32) {
	p.Length = l
}
