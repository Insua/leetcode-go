package main

import "fmt"

func romanToInt(s string) int {
	m := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	num := 0
	for {
		if len(s) == 0 {
			break
		}
		ss := ""
		if len(s) > 1 {
			ss = s[0:2]
		}
		if ss == "CM" || ss == "CD" || ss == "XC" || ss == "XL" || ss == "IX" || ss == "IV" {
			num += m[ss]
			s = s[2:]
		} else {
			ss = s[0:1]
			num += m[ss]
			s = s[1:]
		}
	}
	return num
}

func main() {
	s := "IV"
	fmt.Println(romanToInt(s))
}
