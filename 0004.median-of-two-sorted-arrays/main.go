package main

import "fmt"

// TODO

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k := (m + n) / 2
	fmt.Println(m, n, k, k/2-1)
	return 0
}

func main() {
	nums1 := []int{1, 3, 4, 9}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	findMedianSortedArrays(nums1, nums2)
}
