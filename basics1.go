package main

import "fmt"

// you can omit the type when 2 / > consecutive paramaters share a type
func add(x, y int) int {
	return x + y
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4
	y = sum - x
	// this is known as a naked return
	return
}

// a function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	var a string = "hello"
	var b string = "world"
	a, b = swap(a, b)
	fmt.Println(a, b)
	// variables declared without an explicit initial value are given
	// their zero value (0, false, "")
	var i int
	// variable's type is inferred from the value on the right hand side
	v := 42.2
}
