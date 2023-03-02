package main

import "fmt"

func printBin(num float64) string {
	bin := "0."
	for {
		if num == 0 {
			break
		}
		num *= 2
		if num >= 1 {
			bin += "1"
			num -= 1
		} else {
			bin += "0"
		}
		if len(bin) > 32 {
			bin = "ERROR"
			break
		}
	}
	return bin
}

func main() {
	num := 0.1
	fmt.Println(printBin(num))
}
