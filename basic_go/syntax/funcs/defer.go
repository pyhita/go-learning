package main

// 1 多个defer 之间的执行顺序
func DeferOrder() {
	defer println("1. defer")

	defer println("2. defer")

	defer println("3. defer")
}

// 2 defer 和 return 先后顺序问题
func TestDeferAndReturn() int {
	defer deferCall()

	return returnCall()
}

func returnCall() int {
	println("call return ... ")
	return 1
}

func deferCall() {
	println("call defer ...")
}

// 3 defer and closure
// 调用的时候来确定 i的值，所以是1
func DeferClosure1() {
	i := 0
	defer func() {
		println("i = ", i)
	}()
	i = 1
}

// 传递参数的时候已经确定i的值了，所以是0
func DeferClosure2() {
	i := 0
	defer func(v int) {
		println("i = ", i)
	}(i)
	i = 1
}

// 4 defer 可以修改返回值
// 将返回值copy到另一块内存
func DeferReturnV1() int {
	i := 0
	defer func() {
		i = 1
	}()
	return i
}

// 自始至终都是同一块内存
func DeferReturnV2() (i int) {
	i = 0
	defer func() {
		i = 1
	}()
	return i
}

type Person struct {
	Name string
}

// 指针指向的是同一块内存
func DeferReturnV3() *Person {
	person := &Person{
		Name: "kante",
	}

	defer func() {
		person.Name = "liam"
	}()

	return person
}

// defer 的几个测试demo
func TestDefer1() {
	println("test demo 1 .......")
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func TestDefer2() {
	println("test demo 2 ......")
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func TestDefer3() {
	println("test demo 3 ......")
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
