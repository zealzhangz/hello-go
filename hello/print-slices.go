package hello

import "fmt"

func PrintSlice(s string, x [] int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}