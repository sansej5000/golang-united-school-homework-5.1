package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (c Square) Area() uint {
	return c.a * c.a
}

func (c Square) Perimeter() uint {
	return 4 * c.a
}

func (c Square) End() Point {
	return Point{c.start.x + int(c.a), c.start.y + int(c.a)}
}
