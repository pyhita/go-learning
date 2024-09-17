package main

func main() {
	// 当我们拿到fn 的时候 Closure已经调用完成了
	// 但是我们在执行fn的时候 依然需要Closure中的name
	//fn := Closure1("kante")
	//println(fn())

	//getAge := Closure2()
	//println(getAge())
	//println(getAge())
	//println(getAge())
	//
	//getAge = Closure2()
	//println(getAge())
	//println(getAge())
	//println(getAge())

	//DeferOrder()
	//TestDeferAndReturn()

	//println(DeferReturnV1())
	//println(DeferReturnV2())
	//println(DeferReturnV3().Name)

	// 测试defer 和 返回值
	TestDefer1()
	TestDefer2()
	TestDefer3()
}

// 函数的四种返回值形式
func foo1() int {
	return 2
}

func foo2() (int, int) {
	return 1, 2
}

func foo3() (a, b int) {
	a = 1
	b = 1
	return
}

func foo4() (a int, b string) {
	a = 2
	b = "hello go"
	return
}
