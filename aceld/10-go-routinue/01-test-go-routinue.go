package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0

	for {
		i++
		fmt.Println("child go routinue ", i)
		time.Sleep(1 * time.Second)
	}
}


func main() {
	// 创建一个新的go routinue
	go newTask()

	// main go routinue
	i := 0
	for {
		i++
		fmt.Println("main go routinue, ", i)
		time.Sleep(1 * time.Second)
	}

}