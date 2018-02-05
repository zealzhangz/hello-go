package main

import (
	"fmt"
	"hello-go/hello"
)

var m1 = map[string]hello.VertexLL{
	"Bell Labs":hello.VertexLL{40.68433, -74.39967},
	"Google":hello.VertexLL{37.42202, -122.08408,},
}

func main() {
	fmt.Println(m1)
}
