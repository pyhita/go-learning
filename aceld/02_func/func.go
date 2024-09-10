package main

import (
	"fmt"
)

// 第一种：一个返回值
func foo1(a int, b string) int {
	fmt.Println("--foo1--")

	return a * 2
}

// 第二种：多个返回值匿名形式
func foo2(a int, b string) (int, int) {
	fmt.Println("--foo2--")

	return a * 3, a * 4
}

// 第三种：多个返回值命名形式
func foo3(a int, b string) (ret1 int, ret2 int) {
	fmt.Println("--foo3--")
	
	ret1 = 3
	ret2 = 4
	return 
}

// 第四种：多个返回值命名形式
func foo4(a int, b string) (ret1, ret2 int) {
	return 
}

// go 语言中的引用传递
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	f1 := foo1(2, "xxxx")
	fmt.Println("f1 = ", f1)

	f2, f3 := foo2(2, "xxx")
	fmt.Println("f2 = ", f2, "f3 = ", f3)

	f4, f5 := foo3(3, "xxx")
	fmt.Println("f4 = ", f4, "f5 = ", f5)

	a, b := 1, 2
	fmt.Println("before swap a = ", a, " b = ", b)
	swap(&a, &b)
	fmt.Println("after swap a = ", a, " b = ", b)
}
