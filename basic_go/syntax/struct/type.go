package main

import "fmt"

// go 中定义新类型的方式

// 第一种：通过type关键字进行类型定义
type MyInt int

// 基于类型字面值定义新类型
type T1 map[int]int
type T2 []string

func TestType1() {
	// int 和 MyInt 的底层类型是一样的，所以可以相互转型
	var a int = 1
	var b MyInt = 2
	c := a + int(b)
	fmt.Println(c)
}

// 第二种：类型别名
type X1 = int
