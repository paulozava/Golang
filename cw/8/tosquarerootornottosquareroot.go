package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareOrSquareRoot([]int{100, 101, 102, 5, 5, 1, 1}))
}

//SquareOrSquareRoot is cheat
func SquareOrSquareRoot(numbers []int) []int {
	var processedNumbers []int
	for _, number := range numbers {
		floatNumber := math.Sqrt(float64(number))
		roundNumber := math.Round(floatNumber)
		var newNumber float64
		if floatNumber-roundNumber == 0.0 {
			newNumber = floatNumber
		} else {
			newNumber = float64(number * number)
		}
		processedNumbers = append(processedNumbers, int(newNumber))
	}
	return processedNumbers
}
