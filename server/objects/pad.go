package objects

type Pad struct {
	X      float32
	Y      float32
	Length float32
}

func NewPad(x, y float32) *Pad {
	return &Pad{x, y, 0}
}

func (p *Pad) UpdatePadPos(x, y float32) {
	p.X = x
	p.Y = y
}

func (p *Pad) UpdatePadLength(l float32) {
	p.Length = l
}

func (p *Pad) GetPadRange(canvasWidth float32) map[string]float32 {
	var xl, xr float32

	xl = p.X
	if xl > (canvasWidth - p.Length) {
		xl = canvasWidth - p.Length
	}
	xr = xl + p.Length

	return map[string]float32{
		"XLeft":  xl,
		"XRight": xr,
		"Y":      p.Y,
	}
}
