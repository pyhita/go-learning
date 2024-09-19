package main

import (
	"fmt"
)

// 普通的常量声明
func TestConst1() {
	// 单行常量声明
	const Pi = 3.1415926
	fmt.Println(Pi)

	// const 块声明常量
	const (
		a       = 1
		b       = "test review"
		i, j, s = 12, 11, "bbb"
	)

	fmt.Println(i, j, s)
}

// 无类型常量
type myInt int

func TestConst2() {
	//const i myInt = 33
	//var a int = 2
	// 编译会报错，类型不一致
	// fmt.Println(i + a)

	const i = 33
	var a myInt = 2
	fmt.Println(i + a)
}

// iota 和 枚举
func TestConst3() {
	// 自动重复上一行特性
	const (
		Apple, Banana = 11, 22
		Rice, Grape
		Pear, Watermelon
	)

	const (
		_ = iota
		MON
		TUS
		WEN
		THR
		FRI
		SAT
		SUN
	)

	const (
		_ = iota // 0
		Pin1
		Pin2
		Pin3
		_
		Pin5 // 5
	)
}
