package bar

// Square represents a square.
type Square struct {
	Side float64
}

// Area calculates the area of the shape.
func (c *Square) Area() float64 {
	return c.Side * c.Side
}
