package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	oh := make([]int, len(heights), len(heights))
	copy(oh, heights)
	sort.Ints(heights)
	unMatch := 0
	for k, v := range heights {
		if v != oh[k] {
			unMatch++
		}
	}
	return unMatch
}

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(heights))
}
