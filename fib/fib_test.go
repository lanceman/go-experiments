package fib

import "testing"

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
		got,_ := NextFib(c.given)
		if got != c.expected {
			t.Errorf("Nextfib(%d) == %d, expected %d", c.given, got, c.expected)
		}
	}
}
