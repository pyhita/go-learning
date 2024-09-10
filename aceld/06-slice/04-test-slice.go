package main

import "fmt"

// 切片的追加操作

func main() {
	s1 := make([]int, 4, 5)
	fmt.Println("s1 len is", len(s1), "cap is ", cap(s1))

	// append 
	s1 = append(s1, 1)
	fmt.Println("s1 len is", len(s1), "cap is ", cap(s1))

	// append
	s1 = append(s1, 2)
	fmt.Println("s1 len is", len(s1), "cap is ", cap(s1))
}