package calc_test

import (
	"calctest/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	got := calc.Add(3, 2)
	expected := 5

	if got != expected {
		t.Errorf("not expected result, got:  %v, expected: = %v", got, expected)
	}
}

func TestSub(t *testing.T) {
	got := calc.Sub(3, 2)
	expected := 1
	if got != expected {
		t.Errorf("not got expected result, got: = %v, expected: = %v", got, expected)
	}
}
