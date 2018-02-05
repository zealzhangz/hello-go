package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strArray := strings.Fields(s)
	count := make(map[string]int)
	for i := 0;i < len(strArray); i++{
		str := strArray[i]
		_,ok := count[str]
		if !ok{
			count[str] = 1
		} else {
			count[str] += 1
		}
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
