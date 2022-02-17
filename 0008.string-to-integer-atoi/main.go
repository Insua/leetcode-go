package main

import (
	"fmt"
)

func myAtoi(s string) int {
	ps := ""
	hasPrefixZero := false
	for _, v := range s {
		if v == 43 {
			if len(ps) == 0 {
				ps = "+"
			} else {
				break
			}
		} else if v == 32 {
			if hasPrefixZero {
				return 0
			} else if len(ps) > 0 {
				break
			}
		} else if v == 45 {
			if len(ps) == 0 {
				if hasPrefixZero {
					return 0
				}
				ps += "-"
			} else {
				break
			}
		} else if v == 48 {
			if len(ps) == 0 {
				hasPrefixZero = true
			}
			if len(ps) == 1 && ps != "-" {
				ps += "0"
			} else if len(ps) > 1 {
				ps += "0"
			}
		} else if v >= 48 && v <= 57 {
			ps += string(v)
		} else {
			if len(ps) == 0 {
				return 0
			} else {
				break
			}
		}
	}
	if len(ps) > 0 && ps[0:1] == "+" {
		ps = ps[1:]
	}
	if ps == "-" || ps == "+" {
		return 0
	}
	negative := false
	if len(ps) > 0 && ps[0:1] == "-" {
		negative = true
	}
	if negative {
		ps = ps[1:]
	}
	overNumber := func() int {
		if negative {
			return -2147483648
		}
		return 2147483647
	}
	max := "2147483648"
	if len(ps) > 10 {
		return overNumber()
	}
	if len(ps) == 10 {
		for k, v := range ps {
			if k != 9 {
				if int(v) < int(max[k]) {
					break
				}
				if int(v) > int(max[k]) {
					return overNumber()
				}
			} else {
				last := '7'
				if negative {
					last++
				}
				if int(v) > int(last) {
					return overNumber()
				}
			}
		}
	}
	lps := len(ps)
	mapping := make(map[int32]int)
	for i := 0; i < 10; i++ {
		mapping[int32(48+i)] = i
	}

	pow := func(num, exponent int) int {
		if exponent == 0 {
			return num
		}
		ten := 1
		for i := 0; i < exponent; i++ {
			ten = ten * 10
		}
		return num * ten
	}
	num := 0
	for k, v := range ps {
		num += pow(mapping[v], lps-k-1)
	}
	if negative {
		num = 0 - num
	}
	return num
}

func main() {
	s := "    +1146905820n1"
	fmt.Println(myAtoi(s))
}
