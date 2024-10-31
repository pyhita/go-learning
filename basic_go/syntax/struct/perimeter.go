package main

import (
	"math"
)

func Perimeter(width, length float64) float64 {

	return 2 * (width + length)
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Area(rectangle Rectangle) float64 {

	return rectangle.Width * rectangle.Height
}
