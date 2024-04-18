package poi

type point struct {
	X int64
	Y int64
}

func newPoint(x, y int64) *point {
	return &point{X: x, Y: y}
}
