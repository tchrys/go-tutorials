package main

import "fmt"

func main() {
	// channels can be buffered
	// sends to a buffered channel block only when the buffer is full
	// receives block when the buffer is empty
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}