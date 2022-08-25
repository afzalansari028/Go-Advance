package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	expected := 3
	if got != expected {
		t.Errorf("Got: 1 + 2 = %v, expected: 1 + 2 =%v", got, expected)
	}
}
