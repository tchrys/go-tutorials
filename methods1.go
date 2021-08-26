package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// an interface type is defined as a set of method signatures
type Abser interface {
	abs() float64
}

func guessType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The number is %v\n", v)
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

// a method is a function with a special receiver argument
func (v Vertex) abs() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

func (v *Vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

type IPAddr [4]byte
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	v := Vertex { 3, 4 }
	// methods with pointer receivers
	// pointer conversion is done automatically
	v.scale(2)
	fmt.Println(v.abs())

	var a Abser
	// a value of interface type can hold any value that implements
	// those methods
	a = v
	fmt.Println(a.abs())
	// an interface value that holds a nil concrete value is itself
	// non-nil. Calling a method on a nil interface is a run-time error

	// var emptyInterface interface{}
	// an empty interface may hold values of any type
	// emptyInterface = 42
	// they are used by code that handles value of unknown type

	// a type assertion provides access to an interface value's
	// underlying concrete value
	// iValue := emptyInterface.(int)
	// to test whether an interface value holds a specific type
	// iValue, ok := i.(int)

	p := Person { "Arthur", 42 }
	fmt.Println(p)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}