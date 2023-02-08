package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	sub := []string{folder[0]}
	for _, v := range folder[1:] {
		last := sub[len(sub)-1]
		if strings.HasPrefix(v, last) {
			tp := strings.TrimPrefix(v, last)
			if !strings.HasPrefix(tp, "/") {
				sub = append(sub, v)
			}
		} else {
			if v != "/" {
				sub = append(sub, v)
			}
		}
	}

	return sub
}

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	sub := removeSubfolders(folder)
	fmt.Println(sub)
}
