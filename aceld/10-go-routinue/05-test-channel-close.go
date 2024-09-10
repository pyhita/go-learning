package main

import "fmt"


func main() {
	c := make(chan int, 3)

	go func ()  {
		for i := 0;i < 3;i++ {
			fmt.Println(i)
			c <- i
		}

		// 关闭channel
		close(c)
	} ()

	// for {
	// 	if data, ok := <-c; ok {
	// 		fmt.Println("read from channel ", data)
	// 	} else {
	// 		break
	// 	}
	// }

	// 可以配合range进行简化
	for data := range c {
		fmt.Println("read from channel ", data)
	}

	fmt.Println("main go routinue exit")
}