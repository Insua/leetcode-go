package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rawInt := x
	reverseInt := 0
	for {
		remainder := x % 10
		reverseInt = reverseInt*10 + remainder
		x = x / 10
		if x == 0 {
			break
		}
	}
	return reverseInt == rawInt
}

func main() {
	x := 1234
	fmt.Println(isPalindrome(x))
}
