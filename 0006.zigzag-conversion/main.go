package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	period := 2*numRows - 2
	ss := make([]string, numRows)
	for k, v := range s {
		periodNum := (k + 1) % period
		if periodNum == 0 {
			periodNum = period
		}
		if periodNum <= numRows {
			ss[periodNum-1] += string(v)
		} else {
			ss[period+2-periodNum-1] += string(v)
		}
	}
	zs := ""
	for _, v := range ss {
		zs += v
	}
	return zs
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}
