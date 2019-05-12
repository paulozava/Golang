package main

import "fmt"

func main() {
	fmt.Println(PositiveSum([]int{1, -1, -12, 6}))
}

//PositiveSum takes an array of int and return the sum of positive ones
func PositiveSum(numbers []int) int {
	accumulator := 0
	for _, number := range numbers {
		if number > 0 {
			accumulator += number
		}
	}
	return accumulator
}
func combat(health, damage float64) float64 {

}
