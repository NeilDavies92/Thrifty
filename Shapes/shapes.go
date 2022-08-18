package perimiter

import (
	"math"
)

// Interfaces are a very powerful concept in statically typed languages
// like Go because they allow you to make functions that can be used with
// different types and createhighly-decoupled code whilst still maintaining
// type-safety.

type Shape interface {
	Area() float64
}

// Shape structs
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Shape Types with added methods
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
