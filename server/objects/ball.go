package objects

type Ball struct {
	X uint
	Y uint
}

type EncodedBall struct {
	MessageType int                    `json:"t"`
	Data        map[string]interface{} `json:"d"`
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

func (b *Ball) Encode() *EncodedBall {
	return &EncodedBall{
		2,
		map[string]interface{}{
			"x": b.X,
			"y": b.Y,
		},
	}
}
