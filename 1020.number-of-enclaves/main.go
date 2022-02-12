package main

import "fmt"

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	landNum := 0
	through := make(map[string]string)

	var search func(i, j int)

	search = func(i, j int) {
		through[fmt.Sprintf("%v-%v", i, j)] = ""
		next := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for _, v := range next {
			ii := i + v[0]
			jj := j + v[1]
			if ii > 0 && jj > 0 && ii < m-1 && jj < n-1 {
				if grid[ii][jj] == 1 {
					_, ok := through[fmt.Sprintf("%v-%v", ii, jj)]
					if !ok {
						search(ii, jj)
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				landNum++
			}
			if grid[i][j] == 1 && ((i == 0 || i == m-1) || (j == 0 || j == n-1)) {
				search(i, j)
			}
		}
	}
	return landNum - len(through)
}

func main() {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves(grid))
}
