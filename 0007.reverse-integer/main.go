package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	negative := x < 0
	if negative {
		x = 0 - x
	}
	s := strconv.Itoa(x)
	max := "2147483648"
	sr := make([]rune, len(s))
	for i := 0; i < len(s)/2+1; i++ {
		sr[len(s)-i-1] = rune(s[i])
		sr[i] = rune(s[len(s)-i-1])
	}
	s = strings.TrimLeft(string(sr), "0")
	if len(s) > 10 {
		return 0
	}
	if len(s) == 10 {
		for k, v := range s {
			if k != 9 {
				if int(v) < int(max[k]) {
					break
				}
				if int(v) > int(max[k]) {
					return 0
				}
			} else {
				last := '7'
				if negative {
					last++
				}
				if int(v) > int(last) {
					return 0
				}
			}
		}
	}
	x, _ = strconv.Atoi(s)
	if negative {
		return 0 - x
	}
	return x
}

func main() {
	x := -2147483412
	fmt.Println(reverse(x))
}
