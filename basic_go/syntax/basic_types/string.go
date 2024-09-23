package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串定义、输出
func String1() {
	// 会对转义符进行转义
	println("He said: \"she is a monster \"")
	// 所见即所得，转义符也直接打印
	println(`He said: "she is a monster "`)

	var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
	fmt.Println(s)

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

// 字符串组成
func String2() {

	// 字节视角看待string
	var s = "中国人"
	// len 打印的是字符串的字节个数
	fmt.Printf("the length of s = %d \n", len(s))
	// 打印每个字节的内容
	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i])
	}
	fmt.Printf("\n")

	// 字符视角看待 string
	fmt.Println("the character count in s is ", utf8.RuneCountInString(s))
	// 打印每个字符的内容
	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}
	fmt.Printf("\n")
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
