package main

import "fmt"

var count = 0

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	long := 0

	for i := 0; i < len(s); i++ {
		if long > len(s)-i {
			break
		}
		for j := i + 1; j <= len(s); j++ {
			ss := s[i:j]
			if long > len(ss) {
				continue
			}
			if isDuplicate(ss) {
				break
			} else {
				if len(ss) > long {
					long = len(ss)
				}
				continue
			}
		}
	}
	return long
}

func isDuplicate(s string) bool {
	count++
	m := make(map[int32]uint8)
	for _, v := range s {
		m[v] = 1
	}
	return len(s) != len(m)
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
