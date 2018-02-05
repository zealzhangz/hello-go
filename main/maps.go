package main

import (
	"fmt"
	"hello-go/hello"
)
var m map[string] hello.VertexLL

func main() {
	m = make(map[string] hello.VertexLL)
	m["Bell Labs"] = hello.VertexLL{40.68433, -74.39967,}
	fmt.Println(m["Bell Labs"])
}
