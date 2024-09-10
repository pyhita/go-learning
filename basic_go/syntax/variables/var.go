package main

func main() {
	testVar1()
	testVar2()
	TestConst()
}

// 包变量
// 首字母大写，全局可以进行访问
var Global = "Global"

// 首字母小写，只有本包可以进行访问
// 子包也不可以使用
var local = "local"

// 变量的声明方式，局部变量
func testVar1() {
	println("=================")
	// 第一种方式
	var a int = 11
	var b uint = 12
	println(a, b)

	// 第二种方式：类型推断
	var c = 13
	println(c)

	// 声明的同时初始化，类型推断
	d := 14
	println(d)
}

// 多变量声明，局部变量
func testVar2() {
	println("=================")
	var a, b int = 1, 2
	println(a, b)

	var c, d = 1.1, 2.2
	println(c, d)

	e, f := 11, "xxx"
	println(e, f)

	_, x := 1, "xx"
	println(x)
}
