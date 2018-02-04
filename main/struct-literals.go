package main

import (
	"fmt"
	"../hello"
	)

var (
	v1 = hello.Vertex{1,2}
	v2 = hello.Vertex{X:1}
	v3 = hello.Vertex{}
	p = &hello.Vertex{1,2}
)

func main() {
	fmt.Println(v1,p,v2,v3)
}
