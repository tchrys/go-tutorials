package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	sum := 0
	// [init statement] / condition statement / [post statement]
	for i := 0; i < 10; i++ {
		sum += i
	}
	// if statement can start with a short statement to execute
	// before the condition
	if v:= 2; v < 5 {
		fmt.Println("aaa")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 12:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}