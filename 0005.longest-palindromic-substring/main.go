package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	long := s[0:1]
	for i := 0; i < len(s); i++ {
		if len(long) > len(s)-i {
			break
		}
		for j := i + 1; j <= len(s); j++ {
			ss := s[i:j]
			if len(long) > len(ss) {
				continue
			} else if isPalindromic(ss) && len(s) > len(long) {
				long = ss
			}
		}
	}
	return long
}

func isPalindromic(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "abad"
	fmt.Println(longestPalindrome(s))
}
