package main

import (
	"fmt"
	"../hello"
	)



func main() {
	v := hello.Coordinate{2, 6}
	p := &v
	p.X = 1e9

	fmt.Println(v)
}