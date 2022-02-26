package main

import "fmt"

func maximumDifference(nums []int) int {
	l := len(nums)
	max := -1
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			diff := nums[j] - nums[i]
			if diff > max && diff > 0 {
				max = diff
			}
		}
	}
	return max
}

func main() {
	nums := []int{7, 1, 5, 4}
	fmt.Println(maximumDifference(nums))
}
