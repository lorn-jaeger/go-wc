package main

import (
	"os"
	"testing"
)

func TestGetArgZero(t *testing.T) {
	got := GetArgZero()
	want := os.Args[0]

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
