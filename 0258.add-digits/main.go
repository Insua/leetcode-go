package main

import "fmt"

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	nn := 0
	for {
		if num == 0 {
			break
		}
		nn += num % 10
		num = num / 10
	}
	return addDigits(nn)
}

func main() {
	num := 38
	add := addDigits(num)
	fmt.Println(add)
}
