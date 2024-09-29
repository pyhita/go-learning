package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
)

// 1 透明的错误策略
func HnadleErr1() error {
	err := errors.New("this is an error")

	if err != nil {
		// 打印一些log，然后直接中断处理
		return err
	}

	return nil
}

// 2 哨兵错误处理逻辑

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

func HnadleErr2() error {
	r := bufio.NewReader(bytes.NewBufferString("Hello, World!"))
	data, err := r.Peek(5)

	if err != nil {
		switch err {
		case ErrBufferFull:
			return ErrBufferFull
		case ErrInvalidUnreadByte:
			return ErrInvalidUnreadByte
		case ErrInvalidUnreadRune:
			return ErrInvalidUnreadRune
		}
	}

	// 通过isError 来判断error是不是某个哨兵错误值
	println(errors.Is(err, ErrInvalidUnreadByte))
	println(errors.Is(err, ErrInvalidUnreadRune))

	println(data)
	return nil
}

// 3 自定义错误类型，提供更多的上下文信息
type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

func HandleErr3() {
	err := &MyError{"this is my error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err: %w", err1)

	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2")
		println(e == err)
		return
	}
}

// 4
