package main

import (
	"time"
	"unsafe"
)

// 测试空结构体
func TestEmpty() {
	type Empty struct{}

	var empty1 Empty
	println(unsafe.Sizeof(empty1))

	chan1 := make(chan Empty)
	// 发送一个同步信号
	chan1 <- empty1
}

// 定义一个结构体
type Book struct {
	Title  string
	Price  float64
	Author Person
	Date   time.Time
}

// 结构体之间的嵌套、组合
type Person struct {
	Name string
	Sex  string
	Age  int
}
