package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var arr []int
	for k, v := range nums {
		for kk, vv := range nums[k+1:] {
			if v + vv == target {
				arr = append(arr, []int{k, kk+k+1}...)
			}
		}
	}
	return arr
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
