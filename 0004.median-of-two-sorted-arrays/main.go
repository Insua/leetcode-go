package main

import "fmt"

// TODO

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k := (m + n) / 2
	//fmt.Println(k, k+1)
	//fmt.Println(findKth(nums1, nums2, k))
	fmt.Println(findKth(nums1, nums2, k+1))

	return 0
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if k == 1 {
		n10 := nums1[0]
		n20 := nums2[0]
		if n10 < n20 {
			return n10
		}
		return n20
	}

	removeNum := k / 2
	index := removeNum - 1

	if len(nums1) <= removeNum {
		fmt.Println(nums1, nums2, k, removeNum)
		return nums2[k-1]
	}
	if len(nums2) <= removeNum {
		return nums1[k-1]
	}

	if nums1[index] < nums2[index] {
		nums1 = nums1[removeNum:]
	} else {
		nums2 = nums2[removeNum:]
	}

	return findKth(nums1, nums2, k-removeNum)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10}
	findMedianSortedArrays(nums1, nums2)
}
