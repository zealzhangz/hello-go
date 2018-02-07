package main

import (
	"math"
	"fmt"
)

type VertexTest1 struct {
	X, Y float64
}

func  Abs(v VertexTest1)float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main()  {
	v := VertexTest1{3,4}
	fmt.Println(Abs(v))
}