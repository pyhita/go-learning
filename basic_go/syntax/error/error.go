package main

import (
	"errors"
	"fmt"
	"net"
)

// go 语言中error的类型
func CreateError() {
	err1 := errors.New("first error")
	err2 := fmt.Errorf("index %d is out of bounds", 3)

	fmt.Println(err1, err2)
}

// 自定义error 类型
type OpError struct {
	Op     string
	Net    string
	Source net.Addr
	Addr   net.Addr
	Err    error
}

func (e *OpError) Error() string {
	return e.Op
}
