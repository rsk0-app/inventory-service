package main

import "testing"

func TestAvailable(t *testing.T) {
	cases := []struct {
		stock, reserved, want int
	}{
		{10, 3, 7},
		{5, 5, 0},
		{5, 8, 0},
		{0, 0, 0},
	}
	for _, c := range cases {
		if got := Available(c.stock, c.reserved); got != c.want {
			t.Fatalf("Available(%d,%d) = %d, want %d", c.stock, c.reserved, got, c.want)
		}
	}
}
