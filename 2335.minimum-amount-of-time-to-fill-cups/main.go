package main

import (
	"fmt"
	"sort"
)

func fillCups(amount []int) int {
	count := 0

	var rec func()
	rec = func() {
		if amount[0]+amount[1]+amount[2] == 0 {
			return
		}
		sort.Ints(amount)
		first := amount[0]
		amount[0] = amount[2]
		amount[2] = first
		amount[0] = amount[0] - 1
		if amount[1] > 0 {
			amount[1] = amount[1] - 1
		}
		count++
		rec()
	}

	rec()

	return count
}

func main() {
	amount := []int{5, 0, 0}
	min := fillCups(amount)
	fmt.Println(min)
}
