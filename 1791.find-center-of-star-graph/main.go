package main

import "fmt"

func findCenter(edges [][]int) int {
	for _, v := range edges[0] {
		for _, vv := range edges[1] {
			if v == vv {
				return v
			}
		}
	}
	return 0
}

func main() {
	edges := [][]int{{1, 2}, {2, 3}, {4, 2}}
	fmt.Println(findCenter(edges))
}
