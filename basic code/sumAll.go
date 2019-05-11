package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args
	acc := 0
	for index := 0; index < len(numbers); index++ {
		number, _ := strconv.Atoi(numbers[index])
		acc += number
	}
	fmt.Println(acc)
}
