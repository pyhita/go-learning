package main

import (
	"fmt"
	"runtime"
	"time"
)


func main() {
	// 使用go 创建一个匿名函数
	go func() {
		defer fmt.Println("A defer.")

		// 创建内部匿名函数
		func ()  {
			defer fmt.Println("B defer.")
			// 退出当前的协程
			runtime.Goexit()
			fmt.Println("B")
		} ()

		fmt.Println("A")
	}()

	time.Sleep(1 * time.Second)
}