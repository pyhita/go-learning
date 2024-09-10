package main

import "fmt"

func change(arr []int) {
	arr[0] = 10
}

func main3() {
	arr := []int{1, 2, 3, 4}

	fmt.Println("arr[0] = ", arr[0])
	change(arr)
	fmt.Println("arr[0] = ", arr[0])	

	// 判断slice是否为空
	if arr == nil {
		fmt.Println("arr is empty")
	}
}
