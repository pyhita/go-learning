package main

import (
	"fmt"
	"time"
)


func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("sub go routinue exit")

		for i := 0;i < 4;i++ {
			c <- i
			fmt.Println("向channel 写入 ", i, "此时channel len 是 ", len(c), "channel cap 是 ", cap(c))
		}
	} ()

	time.Sleep(2 * time.Second)

	for i := 0;i < 3;i++ {
		num := <-c
		fmt.Println("从channel读出 ", num)
	}

	fmt.Println("main go routinue exit")
}
