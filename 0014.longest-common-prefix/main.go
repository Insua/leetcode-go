package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	shortLen := len(strs[0])
	shortKey := 0
	for k, v := range strs {
		if len(v) < shortLen {
			shortLen = len(v)
			shortKey = k
		}
	}
	short := strs[shortKey]
	var find func()
	find = func() {
		l := len(short)
		if l == 0 {
			return
		}
		allHave := true
		for _, v := range strs {
			if v[0:l] != short {
				allHave = false
				break
			}
		}
		if allHave {
			return
		}
		short = short[:l-1]
		find()
	}
	find()
	return short
}

func main() {
	strs := []string{"dog", "racecar", "car"}
	s := longestCommonPrefix(strs)
	fmt.Println(s)
}
