package objects

import "padpad/server/message"
import "fmt"

type direction int

const (
	UP    direction = -1
	DOWN  direction = 1
	LEFT  direction = -1
	RIGHT direction = 1
)

type Ball struct {
	X            float32
	Y            float32
	CanvasHeight int
	CanvasWidth  int
	DirX         direction
	DirY         direction
	AngleX       int
	AngleY       int
	Speed        float32
}

func NewBall(startx, starty float32, h, w int) *Ball {
	return &Ball{
		startx,
		starty,
		h,
		w,
		1,
		1,
		1,
		1,
		1,
	}
}

func (b *Ball) Update() {
	/*	if ((padpos.y >= (Math.round(b.Y) + 4) && padpos.y <= (Math.round(this.y) + 7)) && b.X >= padpos.xLeft && b.X <= padpos.xRight) {
	    b.DirY = -1;
	    b.Speed += 0.05;
	}*/
	if b.X >= float32(b.CanvasWidth-7) {
		b.DirX = LEFT
		b.Speed += 0.05
	}
	if b.X <= 7 {
		b.DirX = RIGHT
		b.Speed += 0.05
	}
	if b.Y >= float32(b.CanvasHeight-7) {
		fmt.Println("Lecę do góry", b.Y, ">=", float32(b.CanvasHeight-7), "(", b.CanvasHeight, "-7 )")
		b.DirY = UP
		b.Speed += 0.05
	}
	if b.Y <= 7 {
		b.DirY = DOWN
		b.Speed += 0.05
	}

	b.X = b.X + float32(b.AngleX*int(b.DirX))*b.Speed
	b.Y = b.Y + float32(b.AngleY*int(b.DirY))*b.Speed

	if b.X < 7 {
		b.X = 7
	}
	if b.Y < 7 {
		b.Y = 7
	}
	if b.X > float32(b.CanvasWidth-7) {
		b.X = float32(b.CanvasWidth - 7)
	}
	if b.Y > float32(b.CanvasHeight-7) {
		b.Y = float32(b.CanvasHeight - 7)
	}
}

func (b *Ball) Encode() *message.Message {
	return &message.Message{
		message.BALL_POSITION_TYPE,
		map[string]float32{
			"x": b.X,
			"y": b.Y,
		},
	}
}
