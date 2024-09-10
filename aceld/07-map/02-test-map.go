package main

import "fmt"

func main() {
	m := make(map[int]int, 10)

	// 添加
	m[1] = 1
	m[2] = 2
	m[3] = 3

	// 遍历
	for k, v := range m {
		fmt.Println("k: ", k, "v: ", v)
	}

	// 删除
	delete(m, 3)

	// 修改
	m[2] = 3
	for k, v := range m {
		fmt.Println("k: ", k, "v: ", v)
	}
}