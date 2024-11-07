package main

import "testing"

func TestEmptyInterface(t *testing.T) {

	t.Run("test empty interface", func(t *testing.T) {
		EmptyInterface()
	})
}

func TestNilInterface(t *testing.T) {

	t.Run("test nil interface", func(t *testing.T) {
		NilInterface()
	})
}
