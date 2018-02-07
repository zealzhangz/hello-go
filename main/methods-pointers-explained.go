package main

import (
	"fmt"
	"math"
)

type VertexTest3 struct{
	X,Y float64
}

func Abs(v VertexTest3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v * VertexTest3,f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexTest3{3,4}
	Scale(&v,10)
	fmt.Println(Abs(v))
}
