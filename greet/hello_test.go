package greet

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Gowtham", "English")
		want := "Hello, Gowtham"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty name is given", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("greet in Tamil", func(t *testing.T) {
		got := Hello("Raja", "Tamil")
		want := "வணக்கம், Raja"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
