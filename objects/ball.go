package objects

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
