package main

import (
	"fmt"
	"sort"
)

func alertNames(keyName []string, keyTime []string) []string {
	tm := make(map[string][]int)
	for k, v := range keyName {
		kt := keyTime[k]
		hour := int(kt[0]-'0')*10 + int(kt[1]-'0')
		minute := int(kt[3]-'0')*10 + int(kt[4]-'0')
		hm := hour*60 + minute
		tmGet, ok := tm[v]
		if ok {
			tm[v] = append(tmGet, hm)
		} else {
			tm[v] = []int{hm}
		}
	}

	names := make([]string, 0, len(tm))

	for k, v := range tm {
		alert := false
		sort.Ints(v)
		for i := 0; i <= len(v)-3; i++ {
			tr := v[i : i+3]
			begin := tr[0]
			end := tr[2]
			if end-begin <= 60 && end-begin > 0 {
				alert = true
				break
			}
		}
		if alert {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	return names
}

func main() {
	keyName := []string{"a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b"}
	keyTime := []string{"04:48", "23:53", "06:36", "07:45", "12:16", "00:52", "10:59", "17:16", "00:36", "01:26", "22:42"}

	names := alertNames(keyName, keyTime)
	fmt.Println(names)
}
