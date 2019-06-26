package perimeter

import "math"

// Shape is any 2d shape
type Shape interface {
	Area() float64
}

// Rectangle is a shape dude
type Rectangle struct {
	Width  float64
	Height float64
}

// Area area fn for rectangles
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is another shape
type Circle struct {
	Radius float64
}

// Area area fn for circles
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a shape
func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (rectangle.Width + rectangle.Height)
	return
}

// Area calculates the area of a shape
func Area(rectangle Rectangle) (area float64) {
	area = rectangle.Width * rectangle.Height
	return
}
