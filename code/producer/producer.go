package main

/* producer-consumer problem in Go */

import (
	"fmt"
)

var done = make(chan bool)

// PROD OMIT
func produce(n int) <-chan int {
	writeChan := make(chan int)
	go func() {
		defer fmt.Println("Producer terminated")
		for i := 0; i < n; i++ {
			writeChan <- i
		}
		done <- true

	}()
	return writeChan
}

// END OMIT

// CONS OMIT
func consume(readChan1, readChan2 <-chan int) {
	for {
		select {
		case value := <-readChan1:
			fmt.Printf("From ch1 %d\n", value)
		case value := <-readChan2:
			fmt.Printf("\tFrom ch2 %d\n", value)
		}
	}
}

// END OMIT

func main() {
	ch1 := produce(5)
	ch2 := produce(5)
	go consume(ch1, ch2)
	<-done
	<-done
}
