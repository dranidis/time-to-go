package main

import "fmt"

func workOnSlice(slice []int) {
	fmt.Printf("copy: %T, %v\n%p %p\n", slice, slice, &slice, &slice[0])
	slice[2] = 99
	// slice = append(slice, 100, 101)
}

func workOnSliceRef(slice *[]int) {
	fmt.Printf("refs: %T, %v\n%p %p\n", slice, slice, slice, &(*slice)[0])
	(*slice)[2] = 999
	*slice = append(*slice, 100, 101)
}

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("main: %T, %v\n%p %p\n", slice, slice, &slice, &slice[0])
	fmt.Println(slice)
	workOnSlice(slice)
	fmt.Println(slice)
	// workOnSliceRef(&slice)
	// fmt.Println(slice)
}

// END OMIT
