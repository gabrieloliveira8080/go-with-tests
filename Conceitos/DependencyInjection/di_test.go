package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chriss")

	got := buffer.String()
	want := "Hello, Chriss"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
