package main

import (
	"fmt"
	"slices"
)

func TestSlice1() {
	s1 := []int{1, 2, 3, 4}
	// s2 是s1 子切片
	s2 := s1[2:]

	fmt.Printf("s1: %v, len %d, cap %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap %d \n", s2, len(s2), cap(s2))

	// s2 扩容之后，s1和s2 不再共享空间
	s2 = append(s2, 5)
	fmt.Printf("s1: %v, len %d, cap %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len %d, cap %d \n", s2, len(s2), cap(s2))
}

func TestSlice2() {

	var s1 []int
	fmt.Println("s1 ", s1, " s1 == nil ", s1 == nil, "len(s1) == 0 ", len(s1) == 0)

	// 分配空间
	s1 = make([]int, 3)
	s1[0] = 0
	s1[1] = 1
	s1[2] = 2
	fmt.Println("s1 ", s1, " s1 == nil ", s1 == nil, "len(s1) ", len(s1), "cap(s1) ", cap(s1))

	// append 操作
	s1 = append(s1, 3, 4, 5)
	fmt.Println("s1 ", s1, " s1 == nil ", s1 == nil, "len(s1) ", len(s1), "cap(s1) ", cap(s1))

	// copy
	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Println("copy s2 ", s2, "len(s2) ", len(s2), "cap(s2) ", cap(s2))
	fmt.Printf("s1 addr %p, s2 addr %p \n", &s1, &s2)

	s3 := []string{"a", "b", "c"}
	s4 := []string{"a", "b", "c"}
	if slices.Equal(s3, s4) {
		fmt.Println("equals")
	}
}
