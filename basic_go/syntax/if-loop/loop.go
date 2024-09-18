package main

import "fmt"

func TestFor() {

	// 类似while for循环
	count := 0
	for count < 10 {
		println("count = ", count)
		count++
	}

	// for 遍历slice
	fmt.Println("for slice")
	s1 := []int{1, 2, 3, 4, 5}
	for i, v := range s1 {
		println("i = ", i, "v = ", v)
	}

	fmt.Println("for map")
	// for 遍历map
	m1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	for k, v := range m1 {
		println("k = ", k, "v = ", v)
	}
}
