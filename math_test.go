package main

import "testing"

func TestSumar(t *testing.T) {
	got := Sumar(2, 3)
	want := 5

	if got != want {
		t.Fatalf("Sumar(2, 3) = %d; want %d", got, want)
	}
}