package main

import "fmt"

func main() {
	// Declaration of a pointer to an int
	var pointerToInt *int // HL

	i := 10

	// &i is the memory address of i
	pointerToInt = &i // HL

	fmt.Printf("%T, %v, %v\n", pointerToInt, pointerToInt, *pointerToInt)

	// *pointerToInt is the content of the pointer
	*pointerToInt = 99 // HL

	fmt.Printf("%T, %v, %v\n", pointerToInt, pointerToInt, *pointerToInt)

	fmt.Println(i)
	// /END OMIT
}
