package main

import "fmt"

type CoordinateFiled struct {
	X int
	Y int
}

func main() {
	v := CoordinateFiled{1,6}
	v.X =4
	fmt.Println(v.X)
}