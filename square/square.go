package main

import "fmt"

//START OMIT
// Square OMIT
type Square struct {
	x float64
}

// All user-defined types can have functions attached (methods)
func (s Square) Area() float64 {
	return s.x * s.x
}

func (s *Square) SetX(x float64) {
	// ps := &s OMIT
	// fmt.Printf("%T, %p, %v\n", ps, ps, *ps) OMIT
	s.x = x
	// fmt.Printf("%T, %p, %v\n", ps, ps, *ps) OMIT
}

func main() {
	// Variables of user-defined types (objects)
	s := Square{5}
	// Calling a method
	fmt.Println(s.Area())
	// ps := &s OMIT
	// fmt.Printf("%T, %p, %v\n", ps, ps, *ps) OMIT
	s.SetX(10)
	fmt.Println(s.Area())
}

// END OMIT
