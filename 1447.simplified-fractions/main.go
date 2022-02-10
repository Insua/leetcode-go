package main

import "fmt"

func simplifiedFractions(n int) []string {
	var nums []string
	if n == 1 {
		return nums
	}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if !isGcdBigThan1(i, j) {
				nums = append(nums, fmt.Sprintf("%v/%v", j, i))
			}
		}
	}
	return nums
}

func isGcdBigThan1(a, b int) bool {
	if a == b {
		return false
	}
	if a > b {
		a, b = b, a
	}
	for i := 2; i <= a; i++ {
		if a%i == 0 && b%i == 0 {
			return true
		}
	}
	return false
}

func main() {
	n := 6
	fmt.Println(simplifiedFractions(n))
}
