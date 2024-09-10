package main

import (
	"fmt"
)

// 1 defer 的执行顺序

func a() {
	fmt.Println("a")
}

func b() {
	fmt.Println("b")
}

func c() {
	fmt.Println("c")
}

func main2() {
	defer a()
	defer b()
	defer c()
}