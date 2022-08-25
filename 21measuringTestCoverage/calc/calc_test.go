package calc_test

import (
	"calctest/calc"
	"testing"
)

func TestSubTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 1},
		{x: 2, y: 2, expected: 0},
		{x: 2, y: 3, expected: -1},
	}

	for _, val := range data {
		got := calc.Sub(val.x, val.y)
		if got != val.expected {
			t.Errorf("for x: %v and y: %v expected: %v but got: %v", val.x, val.y, val.expected, got)
		}
	}
}
func TestAddTable(t *testing.T) {
	data := []struct {
		x        int
		y        int
		expected int
	}{
		{x: 2, y: 1, expected: 3},
		{x: 2, y: 2, expected: 4},
		{x: 2, y: 3, expected: 5},
	}

	for _, val := range data {
		got := calc.Add(val.x, val.y)
		if got != val.expected {
			t.Errorf("for x: %v and y: %v expected: %v but got: %v", val.x, val.y, val.expected, got)
		}
	}
}
