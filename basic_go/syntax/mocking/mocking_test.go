package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("test spysleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{Calls: 0}
		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!
`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		if sleeper.Calls != 4 {
			t.Errorf("got %d want %d", sleeper.Calls, 4)
		}
	})

	t.Run("test countdown operator sleeper", func(t *testing.T) {
		countdownOperationSpy := &CountdownOperationSpy{}
		Countdown(countdownOperationSpy, countdownOperationSpy)

		got := countdownOperationSpy.Calls
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
