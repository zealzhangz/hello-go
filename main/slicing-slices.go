package main

import "fmt"

func main()  {
	s := []int{3,5,7,9,11,13,15,17}
	fmt.Println("s ==",s)
	fmt.Println("s[1:4] ==",s[1:4])

	//省略下标代表从0开始
	fmt.Println("s[:4] ==",s[:4])

	//省略上标到len(s)结束
	fmt.Println("s[4:] ==",s[4:])
}
