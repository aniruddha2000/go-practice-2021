package gohellotest

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessege := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Want %q but got %q", want, got)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Tom", "")
		want := "Hello, Tom"
		assertCorrectMessege(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessege(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessege(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessege(t, got, want)
	})
}
