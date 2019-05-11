package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args
	var acc float64
	for index := 0; index < len(numbers); index++ {
		number, _ := strconv.ParseFloat(numbers[index], 10)
		acc += number
	}
	mean := acc / float64(len(numbers)-1)
	fmt.Printf("mean: %f\n", mean)
}
