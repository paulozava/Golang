package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(HowManyDalmatians(33))
}

func HowManyDalmatians(number int) string {
	switch {
	case number < 11:
		return "Hardly any"
	case number < 51:
		return "More than a handful!"
	case number < 101:
		return "Woah that's a lot of dogs!"
	case number == 101:
		return "101 DALMATIONS!!!"
	default:
		return strconv.Itoa(number)
	}
}
