package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	ret := make([][]int, 0)
	items1Map := make(map[int]int)
	for _, v := range items1 {
		items1Map[v[0]] = v[1]
	}
	retMap := make(map[int]int)
	retS := make([]int, 0)
	for _, v := range items2 {
		vv := v[1]
		if v1, ok := items1Map[v[0]]; ok {
			vv += v1
			delete(items1Map, v[0])
		}
		retMap[v[0]] = vv
		retS = append(retS, v[0])
	}
	for k, v := range items1Map {
		retMap[k] = v
		retS = append(retS, k)
	}
	sort.Ints(retS)
	for _, v := range retS {
		ret = append(ret, []int{v, retMap[v]})
	}
	return ret
}

func main() {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}

	ret := mergeSimilarItems(items1, items2)
	fmt.Println(ret)
}
