package main

import "fmt"

// 1 多个defer 之间的执行顺序
func DeferOrder() {
	defer println("1. defer")

	defer println("2. defer")

	defer println("3. defer")
}

// 2 defer 和 return
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

// 3 测试那些函数 可以被defer

func bar() (int, int) {
	return 1, 2
}

func TestDefer3() {
	var c chan int
	var sl []int
	var m = make(map[string]int, 10)
	m["item1"] = 1
	m["item2"] = 2
	//var a = complex(1.0, -1.4)

	var sl1 []int
	defer bar()
	//defer append(sl, 11)
	//defer cap(sl)
	defer close(c)
	//defer complex(2, -2)
	defer copy(sl1, sl)
	defer delete(m, "item2")
	//defer imag(a)
	//defer len(sl)
	//defer make([]int, 10)
	//defer new(*int)
	defer panic(1)
	defer print("hello, defer\n")
	defer println("hello, defer")
	//defer real(a)
	defer recover()
}

// 4 defer 和表达式求值
// defer关键字后面的表达式，是在将deferred函数注册到deferred函数栈的时候进行求值的。

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

// 1.22版本中loopvar语义发生了变更，从循环变量i从per loop var变为per iteration var了。
// 所以每次进入循环，闭包捕获的都是一个新变量i，且该新变量i的值为当次迭代时的i值。等价于1.22版本以前的如下代码：
func foo3() {
	for i := 0; i <= 3; i++ {
		i := i
		defer func() {
			fmt.Println(i)
		}()
	}
}

func TestDefer4() {
	fmt.Println("foo1 result:")
	foo1()
	fmt.Println("\nfoo2 result:")
	foo2()
	fmt.Println("\nfoo3 result:")
	foo3()
}

// 5 defer 性能损耗
