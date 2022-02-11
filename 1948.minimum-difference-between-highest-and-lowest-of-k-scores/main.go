package main

import (
	"sort"
)

func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	min := 0
	for i := 0; i <= len(nums)-k; i++ {
		sn := nums[i : k+i]
		small := sn[k-1] - sn[0]
		if i == 0 {
			min = small
		} else if small < min {
			min = small
		}
	}
	return min
}

func main() {
	nums := []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
	k := 6
	minimumDifference(nums, k)
}
