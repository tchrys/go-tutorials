package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

// function closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibbonaci() func() int {
	first := 0
	second := 1
	return func() int {
		third := first + second
		defer fibAux(&first, &second, third)
		return first
	}
}

func fibAux(first *int, second *int, third int) {
	*first = *second
	*second = third
}

func wordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, v:= range strings.Split(s, " ") {
		if _, ok := ans[v]; ok {
			ans[v] = ans[v] + 1
		} else {
			ans[v] = 1
		}
	}
	return ans
}

func main() {
	v := Vertex { 1, 2 }
	v.X = 4
	p := &v
	// it is not mandatory to write (*p).v
	p.X = 10

	v2 := Vertex { X: 1 } // Y:0 is implicit
	fmt.Println(v2)
	// arrays can not be resized
	primes := [6]int {2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// a slice does not store any data, it just describes a section of
	// an underlying array
	// changing the elements of a slice modifies the corresponding
	// elements of its underlying array
	names := [4]string { "John", "Paul", "George", "Ringo" }
	names1 := names[0:2]
	names2 := names[1:3]
	names2[0] = "X"
	fmt.Println(names1, names2, names)

	arrayLiteral := [3]bool {true, true, false}
	fmt.Println(arrayLiteral)
	// this creates the same array, then builds a slice to reference it
	arraySlice := []bool{true, true, false}
	fmt.Println(arraySlice)

	structSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(structSlice)

	// the zero value of a slice / map is nil
	// the make function allocates a zeroed array and returns a slice
	// that refers to that array
	sliceWithMake1 := make([]int, 5) // len = 5
	sliceWithMake2 := make([]int, 0, 5) // len = 0, cap = 5
	sliceWithMake2 = sliceWithMake2[:cap(sliceWithMake2)] // len,cap=5
	fmt.Println(sliceWithMake1, sliceWithMake2)

	sliceOfSlices := [][]string {
		[]string {"_", "_"},
		[]string {"_", "_"},
	}
	fmt.Println(sliceOfSlices)

	sliceForAppend := make([]int, 0)
	sliceForAppend = append(sliceForAppend, 1, 2)
	fmt.Println(sliceForAppend)

	var firstMap map[string]Vertex = make(map[string]Vertex)
	firstMap["Bell"] = Vertex { 1, 2 }
	fmt.Println(firstMap["Bell"])

	// vertex can be omited
	var mapLiteral = map[string]Vertex {
		"Bell": Vertex { 2, 3 },
	}
	fmt.Println(mapLiteral)
	// insert: m[key] = elem
	// retrieve: elem, ok = m[key]; ok is true if key exists
	// if it doesn't exist -> elem has zero value, ok is false
	// delete: delete(m, key)
	fmt.Println(wordCount("ana are mere"))

	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}

	myFib := fibbonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(myFib())
	}
}