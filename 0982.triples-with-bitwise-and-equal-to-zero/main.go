package main

import "fmt"

func countTriplets(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]&nums[j]&nums[k] == 0 {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	nums := []int{2, 1, 3}
	fmt.Println(countTriplets(nums))
}
