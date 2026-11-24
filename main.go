package main

import (
	"fmt"
	"sort"
)

// EraseOverlapIntervals returns the minimum number of intervals to remove
// so the rest don't overlap (touching endpoints are fine). Equivalent to
// keeping the maximum compatible set — the activity-selection greedy —
// and removing everything else.
func EraseOverlapIntervals(intervals [][2]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sorted := append([][2]int{}, intervals...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][1] < sorted[j][1] })
	kept := 1
	lastEnd := sorted[0][1]
	for _, iv := range sorted[1:] {
		if iv[0] >= lastEnd { // half-open semantics: touching is compatible
			kept++
			lastEnd = iv[1]
		}
	}
	return n - kept
}

func main() {
	fmt.Println(EraseOverlapIntervals([][2]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(EraseOverlapIntervals([][2]int{{1, 2}, {1, 2}, {1, 2}}))
}
