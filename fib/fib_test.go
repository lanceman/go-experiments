package fib

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b, expected uint64
	}{
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 0},
		{1, 1, 2},
		{^uint64(0), 1, 0},
		{^uint64(0) - 1, 2, 0},
	}
	for _, c := range cases {
		got := Add(c.a, c.b)
		if got != c.expected {
			t.Errorf("Add(%d,%d) == %d, expected %d", c.a, c.b, got, c.expected)
		}
	}
}

func TestNextFib(t *testing.T) {
	cases := []struct {
		given, expected uint64
	}{
		{0, 1}, {1, 2}, {2, 3}, {3, 5},
		{4, 5}, {5, 8}, {6, 8}, {7, 8},
		{8, 13}, {12, 13}, {14, 21},
		{54, 55},
	}
	for _, c := range cases {
		got := NextFib(c.given)
		if got != c.expected {
			t.Errorf("Nextfib(%d) == %d, expected %d", c.given, got, c.expected)
		}
	}
}