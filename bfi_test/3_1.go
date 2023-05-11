package main

import (
	"fmt"
	"sort"
)

func minDiffPair(arr []int, n int) {
	if n <= 1 {
		return
	}

	sort.Ints(arr)

	minDiff := arr[1] - arr[0]
	for i := 2; i < n; i++ {
		minDiff = min(minDiff, arr[i]-arr[i-1])
	}

	for i := 1; i < n; i++ {
		if arr[i]-arr[i-1] == minDiff {
			fmt.Printf("%d %d\n", arr[i-1], arr[i])
		}
	}
}

func main() {

	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Ints(arr)
	ln := len(arr)
	minDiffPair(arr, ln)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
