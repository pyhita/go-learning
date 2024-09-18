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

// go 中的string 和 rune的关系：https://gobyexample.com/strings-and-runes
// https://go.dev/blog/strings
func StringAndRune() {
	s := "你好,Java"
	// len 打印的是字节数 不是 字符数
	fmt.Println("len(s) = ", len(s))
	// 如果想知道string 中有几个 rune，需要用下面
	fmt.Println("runes of s is ", utf8.RuneCountInString(s))

	// 打出每一个具体的字符
	for idx, r := range s {
		fmt.Println("idx = ", idx, "r = ", r)
	}

	for i, w := 0, 0; i < len(s); i += w {
		value, width := utf8.DecodeRuneInString(s[i:])
		fmt.Println("decode value is ", value)
		w = width
	}
}
