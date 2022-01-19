package gostructsmethodsinterfacetest

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Height * t.Base
}
