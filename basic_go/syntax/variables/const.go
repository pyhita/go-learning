package main

import (
	"fmt"
	"unsafe"
)

func TestConst() {
	const (
		CategoryBooks    = iota // 0
		CategoryHealth          // 1
		CategoryClothing        // 2
	)

	const (
		RED   = 0
		GREEN = 1
		BLUE  = 3
	)

	const (
		x = (iota * 2) + 1
		y
		z
	)

	// iota 逐行进行累加
	const (
		a1, b1 = iota*2 + 1, iota*2 + 2 // 1 2
		c1, d1                          // 3 4

		e1, f1 = iota * 2, iota * 3 // 4, 6
	)

	const (
		a = "xxxx"
		b = len(a)
		c = unsafe.Sizeof(a)
	)

	fmt.Println(CategoryBooks)
	fmt.Println(RED)
	fmt.Println(y)
	fmt.Println(a, b, c)
}
