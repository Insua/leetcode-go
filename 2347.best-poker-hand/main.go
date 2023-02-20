package main

import (
	"fmt"
	"sort"
)

func bestHand(ranks []int, suits []byte) string {
	if suits[0] == suits[1] && suits[0] == suits[2] && suits[0] == suits[3] && suits[0] == suits[4] {
		return "Flush"
	}

	rankMap := make(map[int]int)
	for _, v := range ranks {
		if vv, ok := rankMap[v]; ok {
			rankMap[v] = vv + 1
		} else {
			rankMap[v] = 1
		}
	}
	maxSame := 0
	for _, v := range rankMap {
		if v > maxSame {
			maxSame = v
		}
	}
	if maxSame >= 3 {
		return "Three of a Kind"
	}
	if maxSame >= 2 {
		return "Pair"
	}

	sort.Ints(ranks)
	if ranks[0] < ranks[1] && ranks[1] < ranks[2] && ranks[2] < ranks[3] && ranks[3] < ranks[4] {
		return "High Card"
	}

	return ""
}

func main() {
	ranks := []int{5, 6, 2, 13, 1}
	suits := []byte{'d', 'a', 'a', 'b', 'c'}

	t := bestHand(ranks, suits)
	fmt.Println(t)
}
