package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// print1()
	// print2()
	print3()
}

// 单变量声明
func print1() {
	var a int
	fmt.Printf("a = %d \n", a)

	var b int = 10
	fmt.Printf("b = %d \n", b)

	var c = 10
	fmt.Printf("c = %d \n", c)

	d := 10
	fmt.Printf("d = %d \n", d)
}

// 多变量声明
func print2() {
	var a, b int
	fmt.Printf("a = %d, b = %d \n", a, b)

	var c, d int = 1, 2
	fmt.Printf("c = %d, d = %d \n", c, d)

	var e, f = 3, 4
	fmt.Printf("e = %d, f = %d \n", e, f)

	g, h := 5, 6
	fmt.Printf("g = %d, h = %d \n", g, h)

	_, v := 11, 11
	fmt.Printf("v = %d \n", v)
}

const (
    CategoryBooks = iota // 0
    CategoryHealth       // 1
    CategoryClothing     // 2
)

const (
	RED = 0
	GREEN = 1
	BLUE = 3
)

const (
	x = (iota * 2) + 1
	y
	z
)

// iota 逐行进行累加
const (
	a1, b1 = iota * 2 + 1, iota * 2 + 2 // 1 2
	c1, d1 // 3 4

	e1, f1 = iota * 2, iota * 3 // 4, 6
)

const (
	a = "xxxx"
	b = len(a)
	c = unsafe.Sizeof(a)
)

// 定义常量
func print3() {
	fmt.Println(CategoryBooks)
	fmt.Println(RED)
	fmt.Println(y)
	fmt.Println(a, b, c)
}