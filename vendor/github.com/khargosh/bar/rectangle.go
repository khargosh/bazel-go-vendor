package bar

// Rectangle represents a rectangle.
type Rectangle struct {
	Length  float64
	Breadth float64
}

// Area calculates the area of the shape.
func (c *Rectangle) Area() float64 {
	return c.Length * c.Breadth
}
