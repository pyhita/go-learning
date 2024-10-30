package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collections of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		sum := Sum(numbers)
		want := 10
		if sum != want {
			t.Errorf("sum = %d, want %d", sum, 15)
		}
	})

	t.Run("collections of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		want := 15
		if sum != want {
			t.Errorf("sum = %d, want %d", sum, 15)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collections of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	assertMessage := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Helper()
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("collections of 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		assertMessage(t, got, want)
	})

	t.Run("empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4})
		want := []int{0, 4}

		assertMessage(t, got, want)
	})
}
