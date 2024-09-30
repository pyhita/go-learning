package main

import "fmt"

func Foo() {
	println("call Foo")
	Bar()
	println("Foo exit")
}

func Bar() {
	defer func() {
		// recover from the panic
		if err := recover(); err != nil {
			fmt.Println("recover from the panic: ", err)
		}
	}()

	println("call Bar")
	panic("panic occurs in Bar")
	Zoo()
	println("Bar exit")
}

//func Bar() {
//	println("call Bar")
//	panic("panic occurs in Bar")
//	Zoo()
//	println("Bar exit")
//}

func Zoo() {
	println("call Zoo")
	println("Zoo exit")
}
