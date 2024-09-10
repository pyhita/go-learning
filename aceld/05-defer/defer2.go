package main

import "fmt"

// 2 defer 和 return 那个先被执行

func main() {
	returnAndDefer()
}

func returnAndDefer() int {
	defer deferCall()

	return returnCall()
}

func returnCall() int {
	fmt.Println("return call ... ")
	return 1
}

func deferCall() {
	fmt.Println("defer call ... ")
}