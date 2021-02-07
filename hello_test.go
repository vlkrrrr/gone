package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("jeöfsd sfd fs ", func(t *testing.T) {
		got := Hello("Volker")
		want := "Hello, Volker"

		assertCorrectMessage(t, got, want)
	})
	t.Run("jeöfsd sfd fs2222 ", func(t *testing.T) {
		got := Hello("")
		want := "Hello, wurld"
		assertCorrectMessage(t, got, want)

	})

}
