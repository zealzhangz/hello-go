package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs()float64
}

type MyFloat1 float64

func (f MyFloat1)Abs()float64  {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}

type VertexTest7 struct{
	X,Y float64
}

func (v *VertexTest7)Abs()float64  {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat1(math.Sqrt2)
	v := VertexTest7{3,4}

	a = f //a MyFloat 实现了 Abser
	a = &v //a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v //放开改行注释会报编译错误，因为实现Abs是选择指针作为接受者，给a赋值时必须是指针或引用

	fmt.Println(a.Abs())
}

