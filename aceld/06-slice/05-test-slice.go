package main

import "fmt"

// 切片的截取操作

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1[0:2]
	fmt.Println("s1: ", s1, "s2: ", s2)
	
	// 浅拷贝
	s2[0] = 22
	fmt.Println("s1: ", s1, "s2: ", s2)

	// 深拷贝
	s3 := make([]int, 3)
	copy(s3, s1)
	s3[0] = 11
	fmt.Println("s1: ", s1, "s2: ", s2, "s3: ", s3)
}