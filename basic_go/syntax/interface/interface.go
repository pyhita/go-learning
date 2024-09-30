package main

// 测试空接口
func TestEmptyInterface() {
	var i interface{} = 15
	v, ok := i.(int)
	if ok {
		println("i is int type, ", v)
	}

	i = "hello go"
	_, ok = i.(string)
	if ok {
		println("i is string type, ", ok)
	}

	type T struct {
	}

	var t T
	i = t
	i = &t
	println(i)
}
