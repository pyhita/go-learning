package main

import "fmt"

func main() {
	// 利用channel 实现父子协程之间的通信
	c := make(chan int)

	go func () {
		defer fmt.Println("sub go routinue exit")

		fmt.Println("sub go routinue running")
		c <- 666
	} ()

	num := <-c
	fmt.Println("num : ", num)
	fmt.Println("main go routinue exit")
}