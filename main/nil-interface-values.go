package main

import "fmt"

type I2 interface {
	M()
}

func main() {
	var i I2
	describe2(i)
	i.M()

}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
