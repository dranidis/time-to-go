package main

import "fmt"

func workOnArray(array [3]int) {
	fmt.Printf("func: %T, %v\n&array: %p &array[0]: %p\n", array, array, &array, &array[0])
	array[0] = 99
}

func main() {
	array := [3]int{1, 2, 3}
	fmt.Printf("main: %T, %v\n&array: %p &array[0]: %p\n", array, array, &array, &array[0])
	fmt.Println(array)
	workOnArray(array)
	fmt.Println(array)
}
