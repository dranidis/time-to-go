package main

import (
	"errors"
	"fmt"
)

// ERROR OMIT
type DivisionByZero float64

func (e DivisionByZero) Error() string {
	return fmt.Sprintf("Cannot divide %f by zero.", e)
}

// END OMIT

// DIV OMIT
func Division(i, j float64) (float64, error) {
	if j == 0.0 {
		return i, DivisionByZero(i)
	}
	return i / j, nil
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("square root of negative number is not defined")
	}
	return f, nil
}

// END OMIT
// CHECK OMIT
func main() {
	for div := 2; div > -1; div-- {
		if res, err := Division(3.0, float64(div)); err != nil { // HL
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%f\n", res)
		}
	}
	// HL
	if res, err := Sqrt(-1.0); err != nil { // HL
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%f\n", res)
	}
}

// END OMIT
