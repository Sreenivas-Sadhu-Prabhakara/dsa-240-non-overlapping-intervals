package main

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want int
	}{
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][2]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][2]int{{1, 2}, {2, 3}}, 0}, // touching is not overlapping
		{[][2]int{}, 0},
		{[][2]int{{1, 9}, {2, 3}, {4, 5}, {6, 7}}, 1}, // dump the long one
		{[][2]int{{1, 100}}, 0},
	}
	for _, c := range cases {
		if got := EraseOverlapIntervals(c.in); got != c.want {
			t.Errorf("EraseOverlapIntervals(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
