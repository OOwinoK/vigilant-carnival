package greetings

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World. Welcome!"

	if got != want {
		t.Errorf("Hello() = %q; want %q", got, want)
	} else {
		t.Log("Test passed")
	}
}
