package main

import (
	"fmt"
	"sort"
)

func maxValue(grid [][]int, mi, mj int) int {
	mv := make([]int, 0)
	for i := mi - 1; i <= mi+1; i++ {
		for j := mj - 1; j <= mj+1; j++ {
			mv = append(mv, grid[i][j])
		}
	}
	sort.Ints(mv)
	return mv[8]
}

func largestLocal(grid [][]int) [][]int {
	max := make([][]int, len(grid)-2)
	for i := range max {
		v := make([]int, len(grid)-2)
		for j := range v {
			v[j] = maxValue(grid, i+1, j+1)
		}
		max[i] = v
	}
	return max
}

func main() {
	grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
	fmt.Println(largestLocal(grid))
}
