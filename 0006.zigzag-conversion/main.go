package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	var slice [][]string
	for _, v := range s {
		fmt.Println(string(v))
	}
	fmt.Println(slice)
	return ""
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	convert(s, numRows)
}
