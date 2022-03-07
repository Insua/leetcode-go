package main

import "fmt"

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	isNegative := num < 0
	if isNegative {
		num = 0 - num
	}
	s := ""
	for {
		if num == 0 {
			break
		}
		digit := num % 7
		s = string(byte(digit+48)) + s
		num = num / 7
	}
	if isNegative {
		s = "-" + s
	}
	return s
}

func main() {
	num := -7
	b7 := convertToBase7(num)
	fmt.Println(b7)
}
