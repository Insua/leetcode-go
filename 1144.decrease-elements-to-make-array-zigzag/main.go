package main

import "fmt"

func odd(nums []int) int {
	rn := make([]int, len(nums))
	copy(rn, nums)
	var count int
	for k := range rn {
		if k%2 == 1 {
			if k+1 < len(rn) {
				l := rn[k-1]
				m := rn[k]
				r := rn[k+1]
				if l >= m {
					nl := m - 1
					count += l - nl
					rn[k-1] = nl
				}
				if r >= m {
					nr := m - 1
					count += r - nr
					rn[k+1] = nr
				}
			} else {
				l := rn[k-1]
				m := rn[k]
				if l >= m {
					nl := m - 1
					count += l - nl
					rn[k-1] = nl
				}
			}
		}
	}
	return count
}

func even(nums []int) int {
	rn := make([]int, len(nums))
	copy(rn, nums)
	var count int
	for k := range rn {
		if k%2 == 0 {
			if k == 0 {
				m := rn[k]
				r := rn[k+1]
				if r >= m {
					nr := m - 1
					count += r - nr
					rn[k+1] = nr
				}
			} else if k+1 < len(rn) {
				l := rn[k-1]
				m := rn[k]
				r := rn[k+1]
				if l >= m {
					nl := m - 1
					count += l - nl
					rn[k-1] = nl
				}
				if r >= m {
					nr := m - 1
					count += r - nr
					rn[k+1] = nr
				}
			} else {
				l := rn[k-1]
				m := rn[k]
				if l >= m {
					nl := m - 1
					count += l - nl
					rn[k-1] = nl
				}
			}
		}
	}
	return count
}

func movesToMakeZigzag(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	o := odd(nums)
	e := even(nums)
	if o < e {
		return o
	}
	return e
}

func main() {
	nums := []int{7, 4, 8, 9, 7, 7, 5}
	i := movesToMakeZigzag(nums)
	fmt.Println(i)
}
