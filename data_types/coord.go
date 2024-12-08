package datatypes

import "strconv"

type Coord struct {
	X int
	Y int
}

func NewCoord(x int, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c *Coord) GetDiff(next Coord) (int, int) {
	xDiff := next.X - c.X
	yDiff := next.Y - c.Y

	return xDiff, yDiff
}

func (c *Coord) Add(xDiff int, yDiff int) Coord {
	return Coord{
		X: c.X + xDiff,
		Y: c.Y + yDiff,
	}
}

func (c *Coord) Sub(xDiff int, yDiff int) Coord {
	return Coord{
		X: c.X - xDiff,
		Y: c.Y - yDiff,
	}
}

func (c *Coord) ToString() string {
	return "(" + strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y) + ")"
}
