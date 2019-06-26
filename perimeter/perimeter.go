package perimeter

import "math"

// Shape is any 2d shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle is a shape dude
type Rectangle struct {
	Width  float64
	Height float64
}

// Area area fn for rectangles
func (r Rectangle) Area() (area float64) {
	area = r.Width * r.Height
	return
}

// Perimeter of Rectangle
func (r Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (r.Width + r.Height)
	return
}

// Circle is another shape
type Circle struct {
	Radius float64
}

// Area area fn for circles
func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

// Perimeter of Circle
func (c Circle) Perimeter() (perimeter float64) {
	perimeter = 2 * c.Radius * math.Pi
	return
}

// Triangle is a shape
type Triangle struct {
	Base   float64
	Height float64
}

// Area of Triangle
func (t Triangle) Area() (area float64) {
	area = 0.5 * t.Base * t.Height
	return
}

// Perimeter of Triangle
func (t Triangle) Perimeter() (perimeter float64) {
	perimeter = t.Base + t.Height + math.Hypot(t.Base, t.Height)
	return
}
