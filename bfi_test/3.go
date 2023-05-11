package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	sort.Ints(array)
	minAbs := int(^uint(0) >> 1)
	var pairs string
	for i := 0; i < n-1; i++ {
		var absDiff int
		if (array[i] < 0 && array[i+1] < 0) || (array[i] > 0 && array[i+1] > 0) {
			absDiff = abs(array[i] - array[i+1])
		} else {
			absDiff = abs(array[i]) + abs(array[i+1])
		}
		if absDiff < minAbs {
			minAbs = absDiff
			pairs = fmt.Sprintf("%d %d\n", array[i], array[i+1])
		} else if absDiff == minAbs {
			pairs += fmt.Sprintf("%d %d\n", array[i], array[i+1])
		}
	}
	fmt.Println(pairs)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
