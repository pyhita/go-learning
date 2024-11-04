package mocking

import (
	"fmt"
	"io"
)

const (
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	// 记录 sleep 调用次数
	s.Calls++
}

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := 3; i >= 1; i-- {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintln(writer, "Go!")
}
