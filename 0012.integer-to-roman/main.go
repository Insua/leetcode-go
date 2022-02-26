package main

import "fmt"

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	s := ""
	for i := 0; i < 13; i++ {
		for {
			if num >= nums[i] {
				s += symbols[i]
				num -= nums[i]
			} else {
				break
			}
		}
	}
	return s
}

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}
