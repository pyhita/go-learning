package main

// 函数闭包
func Closure1(name string) func() string {

	return func() string {
		return "hello, " + name
	}
}

func Closure2() func() int {
	age := 0

	return func() int {
		age++
		return age
	}
}
