package dependency_inject

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kante")

	got := buffer.String()
	want := "Hello, Kante!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
