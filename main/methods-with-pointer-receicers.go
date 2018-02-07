package main

import (
	"fmt"
	"math"
)

type VertexTest6 struct{
	X,Y float64
}

func (v *VertexTest6)Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *VertexTest6)Abs()float64  {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main()  {
	v := &VertexTest6{6,8}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(10)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}