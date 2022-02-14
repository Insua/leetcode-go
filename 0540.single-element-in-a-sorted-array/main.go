package main

import (
	"fmt"
)

func singleNonDuplicate(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[l-1] != nums[l-2] {
		return nums[l-1]
	}
	low := 0
	high := l - 1
	for low <= high {
		mid := (low + high) / 2
		left := mid - 1
		right := mid + 1
		oddEven := ""
		if nums[mid] == nums[left] {
			oddEven = fmt.Sprintf("%v%v", left%2, mid%2)
		} else if nums[mid] == nums[right] {
			oddEven = fmt.Sprintf("%v%v", mid%2, left%2)
		} else {
			return nums[mid]
		}
		if oddEven == "10" {
			high = mid - 1
		}
		if oddEven == "01" {
			low = mid + 1
		}
	}
	return 0
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}
