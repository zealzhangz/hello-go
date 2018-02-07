package main

import (
	"fmt"
	"math"
)

type VertexTest5 struct {
	X,Y float64
}

func (v VertexTest5)Abs()float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v VertexTest5)float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	v := VertexTest5{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertexTest5{6,8}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

