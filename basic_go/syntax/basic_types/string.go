package main

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	println("He said: \"she is a monster \"")
	println(`He said: "she is a monster "`)

	// 字符串只可以和字符串拼接
	println("Hello, " + "golang")
	// 编译无法通过
	// println("Hello " + 123)
	// 格式化打印
	println(fmt.Sprintf("%s %d", "kante", 22))

	// 字符串操作函数
	println(len("abc"))
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
}
