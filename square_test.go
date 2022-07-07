package square

import "testing"

func Test_area(t *testing.T) {
	expected := uint(9)
	input := Square{start: Point{x: 2, y: 3}, a: 3}
	res := input.Area()
	if expected != res {
		t.Errorf("invalid value, expected %d", expected)
	}
}

func Test_perimeter(t *testing.T) {
	expected := uint(12)
	input := Square{start: Point{x: 2, y: 3}, a: 3}
	res := input.Perimeter()
	if expected != res {
		t.Errorf("invalid value, expected %d", expected)
	}
}

func Test_end(t *testing.T) {
	expected := Point{x: 5, y: 6}
	input := Square{start: Point{x: 2, y: 3}, a: 3}
	res := input.End()
	if expected != res {
		t.Errorf("invalid value, expected %d", expected)
	}
}
