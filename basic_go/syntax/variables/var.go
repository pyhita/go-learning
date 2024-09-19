package main

// 包变量
// 首字母大写，全局可以进行访问
var Global = "Global"

// 首字母小写，只有本包可以进行访问
// 子包也不可以使用
var local = "local"

// 声明多个包变量，更推荐变量块
var (
	t1 = "t1"
	t2 = "t2"
	x1 = 1
)

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

	// 第四种方式 先声明、后赋值
	var g int
	g = 10
	println(g)
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

// 多变量：变量声明块
func testVar3() {
	var (
		a = 11
		b = 11
		c = "string"
		f = 1.1
	)

	println(a, b, c, f)
}
