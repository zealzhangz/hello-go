package main

import (
	"../hello"
	)

func main() {
	var a []int
	hello.PrintSlice("a",a)

	// append works on nil slices.
	a = append(a,0)
	hello.PrintSlice("a",a)

	// the slice grows as needed.
	a = append(a,1)
	hello.PrintSlice("a",a)

	// we can add more than one element at a time.
	a = append(a,2,3,4,5,6,7,8,)
	hello.PrintSlice("a",a)
}