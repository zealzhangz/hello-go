package main

import "fmt"

type VertexTest4 struct{
	X,Y float64
}

func (v * VertexTest4)Scale(f float64)  {
	v.Y = v.Y * f
	v.X = v.X * f
}

func ScaleFunc(v * VertexTest4,f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}


func main(){
	v := VertexTest4{3,4}
	v.Scale(2)
	ScaleFunc(&v , 10)

	p := &VertexTest4{6,8}
	p.Scale(2)
	ScaleFunc(p,10)

	fmt.Println(v,p)

}
