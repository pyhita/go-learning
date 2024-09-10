package main

import "fmt"

func Byte() {
	var a byte = 'a'
	// 输出的是a的 ASII 码
	println(a)
	// 格式化输出 字节
	println(fmt.Sprintf("%c", a))

	// byte[] 和 string
	var str string = "this a string"
	var chs []byte = []byte(str)
	println(chs)
	// 更改切片
	chs[0] = 'T'
	// 不会影响到原来的str，说明发生了深度拷贝
	println(str)
}
