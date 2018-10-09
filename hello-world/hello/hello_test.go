package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Say hello to a friend", func(t *testing.T) {
		got := Hello("Bangalore")
		want := "Hello, Bangalore"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Empty argument test", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
