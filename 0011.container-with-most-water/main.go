package main

import "fmt"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ma := 0
	for {
		if l >= r {
			break
		}
		lh := height[l]
		rh := height[r]
		area := (r - l) * min(lh, rh)
		if lh > rh {
			r--
		} else {
			l++
		}
		if area > ma {
			ma = area
		}
	}
	return ma
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 1}
	fmt.Println(maxArea(height))
}
