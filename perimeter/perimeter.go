package perimeter

// Rectangle is a shape dude
type Rectangle struct {
	Width  float64
	Height float64
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
