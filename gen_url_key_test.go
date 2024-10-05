package main

import "testing"

func TestGenUrlKey(t *testing.T) {

	got := genUrlKey("https://www.google.com/")
	want := "d0e19"

	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
