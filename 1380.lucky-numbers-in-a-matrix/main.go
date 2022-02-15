package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	var luck []int
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		min := matrix[i][0]
		minJ := 0
		for j := 1; j < n; j++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				minJ = j
			}
		}
		isLuck := true
		for i := 0; i < m; i++ {
			if matrix[i][minJ] > min {
				isLuck = false
				break
			}
		}
		if isLuck {
			luck = append(luck, min)
		}
	}
	return luck
}

func main() {
	matrix := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println(luckyNumbers(matrix))
}
