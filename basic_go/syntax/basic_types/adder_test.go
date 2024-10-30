package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected: %d, got: %d", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(3, 3)
	fmt.Println(sum)
	// Output: 6
}
