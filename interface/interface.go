package main

import "fmt"

type Shape interface {
	SetX(float64)
	Area() float64
}

type Square struct {
	x float64
}

func (s *Square) SetX(x float64) { s.x = x }

func (s Square) Area() float64 { return s.x * s.x }

// INT OMIT
func main() {
	var s Shape
	fmt.Printf("Value/Type of s:  %v:%T\n", s, s)
	var sqPtr *Square
	s = sqPtr
	fmt.Printf("Value/Type of s:  %v:%T\n", s, s)
	s = &Square{4}

	fmt.Printf("Area of %v:%T is %.2f\n", s, s, s.Area())
}

// END OMIT
