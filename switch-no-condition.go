package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	switch {
	case today.Hour() < 12:
		fmt.Println("Good morning.")
	case today.Hour() < 17:
		fmt.Println("Good aftertoon.")
	default:
		fmt.Println("Good evening.")
	}
}
