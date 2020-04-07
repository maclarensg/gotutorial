package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello!"
	if got := Helloworld(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
