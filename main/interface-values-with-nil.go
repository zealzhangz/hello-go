package main

import (
	"fmt"
)

type I1 interface {
	M()
}

type T1 struct {
	S string
}

func (t *T1)M()  {
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I1
	var t *T1
	i = t
	describe1(i)
	i.M()

	i = &T1{"Hello"}
	describe1(i)
	i.M()
}

func describe1(i I1)  {
	fmt.Printf("(%v, %T)\n",i,i)

}


