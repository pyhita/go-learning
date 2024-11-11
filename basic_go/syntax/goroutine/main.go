package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	c := foo(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})

	fmt.Println(<-c, "goroutine exit.")
}

func foo(f func() error) chan error {
	c := make(chan error)
	go func() {
		c <- f()
	}()

	return c
}

func testRebase() {

}
