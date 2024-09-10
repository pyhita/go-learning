package main

import "fmt"

// 1 测试固定长度的静态数组

// 使用静态传参的时候是值拷贝，并且参数的长度必须匹配
func test(arr [4]int) {
	arr[0] = 3
}

func main2() {
	arr1 := [10]int {1, 2, 3, 4}
	arr2 := [4]int {1, 2, 3, 4}

	// 遍历数组
	for i := 0;i < len(arr1);i++ {
		fmt.Println(i, " ", arr1[i])
	}

	for i, r := range arr2 {
		fmt.Println(i, " ", r)
	}

	fmt.Println("arr2[0] = ", arr2[0])
	// test(arr1) 报错，无法通过编译
	test(arr2)
	fmt.Println("arr2[0] = ", arr2[0])
	// 前后两次arr0 的值都是相同的，说明是值传递，不是引用传递
}
