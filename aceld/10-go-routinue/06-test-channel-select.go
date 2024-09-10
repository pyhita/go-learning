package main

import "fmt"

func main() {

	nums := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0;i < 6;i++ {
			// 从nums中读取 斐波那契数
			fmt.Println(<-nums)
		}

		// 退出程序
		quit <- 0
	} ()

	x, y := 1, 1
	for {
		select {
			case nums <- x:
				x = y
				y = x + y
			case <-quit:
				fmt.Println("quit")
				return
		}
	}

	defer fmt.Println("Finished ... ")
}