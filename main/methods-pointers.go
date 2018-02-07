package main

import (
	"fmt"
	"math"
)

type VertexTest2 struct {
	X,Y float64
}

func (v VertexTest2)Abs()float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexTest2)Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexTest2{3,4}
	v.Scale(10)
	fmt.Println(v)
}