package main

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known words", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}

		got, err := dict.Search("test")
		if err != nil {
			t.Fatal("unexpected error")
		}
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown words", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}

		_, got := dict.Search("unknown")
		//want := "could not find the word you were looking for"

		if got == nil {
			t.Fatal("excepted to get an error")
		}

		if got != ErrNotFound {
			t.Errorf("got %v want %v", got, ErrNotFound)
		}
	})
}

func TestAddWord(t *testing.T) {

	t.Run("new words", func(t *testing.T) {
		dict := Dict{}
		dict.Add("test", "this is just a test")
		got, err := dict.Search("test")
		want := "this is just a test"

		if err != nil {
			t.Fatal("unexpected error")
		}

		assertStrings(t, got, want)
	})

	t.Run("existing words", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}
		err := dict.Add("test", "this is just a test")

		if err == nil {
			t.Fatal("expected error but not")
		}

		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update words", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}
		dict.Update("test", "this is just a test v2")
		newWorld := "this is just a test v2"

		assertDesition(t, dict, "test", newWorld)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got.Error(), want.Error())
	}
}

func assertDesition(t *testing.T, dict Dict, key, want string) {
	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("unexpected error")
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
