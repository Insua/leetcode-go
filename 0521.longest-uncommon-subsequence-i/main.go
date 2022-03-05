package main

import "fmt"

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	la := len(a)
	lb := len(b)
	if la > lb {
		return la
	}
	return lb
}

func main() {
	a := "aefawfawfawfaw"
	b := "aefawfeawfwafwaef"
	l := findLUSlength(a, b)
	fmt.Println(l)
}
