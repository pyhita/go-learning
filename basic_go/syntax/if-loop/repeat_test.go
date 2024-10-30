package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected: %s, actual: %s", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	// b.N 一个框架认为的合适的次数
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
