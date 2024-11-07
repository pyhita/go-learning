package main

import "errors"

type MyInterface interface {
	X1()
}

type T int

func (t T) X1() {
	println("T's X1")
}

// 测试接口断言
func EmptyInterface() {

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

	var t T
	var j interface{} = t
	v1, ok := j.(MyInterface)
	if ok {
		v1.X1()
	}
}

// 测试接口 == nil

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func NilInterface() {
	var myError error = nil
	println(myError == nil)

	err := returnsError()
	println(err == nil)
}
