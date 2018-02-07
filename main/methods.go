package main

import (
	"fmt"
	"math"
)

type VertexTest struct {
	X, Y float64
}

func (v VertexTest) Abs()float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := VertexTest{3,4}
	fmt.Println(v.Abs())
}
