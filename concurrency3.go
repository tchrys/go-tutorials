package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x + y
	}
	// a sender can close a channel to indicate that no more
	// values will be sent
	// to test whether a channel has been closed:
	// v, ok := <- ch (ok = false => no more values)
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// receive values from the channel repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}

}