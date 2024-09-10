package main

import "fmt"

// interface{} 空接口是万能类型，可以传参任意类型
func testInterface(arg interface{}) {

	// go语言提供了断言机制来判断arg是哪一种类型
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, the value is ", value)
	} else {
		fmt.Println("arg is not string type")
	}
}

func main() {
	testInterface(1)
	testInterface("test")
}