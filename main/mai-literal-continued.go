package main

import (
	"fmt"
	"hello-go/hello"
)

var m2 = map[string]hello.VertexLL{
	"Bell Labs":{40.68433, -74.39967},
	"Google":{37.42202, -122.08408,},
}

func main() {
	fmt.Println(m2)
}
