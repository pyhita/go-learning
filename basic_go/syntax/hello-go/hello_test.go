package main

import "testing"

func TestHello(t *testing.T) {
	assertMessage := func(t *testing.T, got, want string) {
		// t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）
		// 通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("kante", "engish")
		want := "Hello, kante"
		assertMessage(t, got, want)
	})

	t.Run("empty name", func(t *testing.T) {
		got := Hello("", "engish")
		want := "Hello, World"
		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("kante", "spanish")
		want := "Hola, kante"
		assertMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("kante", "french")
		want := "Bonjour, kante"
		assertMessage(t, got, want)
	})
}
