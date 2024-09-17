package main

import "fmt"

// golang 函数式编程

// 1 函数作为变量进行赋值
func Fun1() int {
	return 100
}

func Fun2() int {
	myFun := Fun1

	return myFun()
}

// 2 局部方法，封装公共逻辑，并且不希望别人使用
func Fun3() string {
	fn := func() {
		println("test ...")
	}
	fn()

	func() {
		println("test2 ...")
	}()

	return "finish"
}

// 3 方法作为返回值
func Fun4() func(a int) string {

	return func(a int) string {
		return fmt.Sprintf("%s : %d", "xxx", a)
	}
}

// 4 不定长参数
func Fun5(name string, ages ...int) int {

	return len(ages)
}
