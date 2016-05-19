package main

import (
	"fmt"

	"github.com/khargosh/bar"
)

type Shape interface {
	Area() float64
}

func printArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	printArea(&bar.Circle{2})
	printArea(&bar.Rectangle{2, 4})
	printArea(&bar.Square{2})
}
