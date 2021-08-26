package main

// a goroutine is a lightweight thread managed by the go runtime
// the evaluation of parameters/function happens in the current
// goroutine, the execution -> next goroutine

// channels are a typed conduit through which you can send and
// receive values with the channel operator, <-
// ch <- v (send v to channel ch)
// v := <-ch (receive from ch, assign value to v)
// by default sends and receives block until the other side is ready
// => synchronize without explicit locks

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sends sum to c
}

func main() {
	s := []int {7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[ : len(s) / 2], c)
	go sum(s[len(s) / 2 : ], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x + y)
}