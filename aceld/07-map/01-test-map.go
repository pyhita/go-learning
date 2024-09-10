package main

import "fmt"

// map 的声明方式

func main() {
	m1 := make(map[string]string, 5)
	if m1 != nil {
		fmt.Println("m1 len: ", len(m1))
	}

	m1["one"] = "java"
	m1["two"] = "c++"
	m1["three"] = "python"

	// 声明并进行初始化
	m2 := map[int]string {
		1 : "java",
		2 : "c++",
		3 : "python",
	}

	fmt.Println("m1: ", m1, "m2: ", m2)
}