package main

import "fmt"

// slice 的四种声明方式

func main4() {
	// 声明的同时进行初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	// 声明，并通过make初始化
	var s2 []int
	// 指定初始容量是5
	s2 = make([]int, 5)
	s2[0] = 33
	fmt.Println("s2 = ", s2)

	// 声明同时通过make分配容量
	s3 := make([]int, 10)
	fmt.Println("s3 = ", s3)

}