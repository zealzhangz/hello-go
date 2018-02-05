package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var array = [][]uint8{}
	for i := 0;i < dy;i++{
		array = append(array, make([]uint8,dx))
	}
	return array
}

func main() {
	pic.Show(Pic)
}