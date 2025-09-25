package main

import (
	"fmt"
	// "math"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (circle Circle) GetCircumference() float64 {
	return 2 * PI * circle.radius // math.Pi
}

func main() {
	circle := Circle{
		radius: 4,
	}

	fmt.Println(circle.GetCircumference())
}
