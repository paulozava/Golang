package main

import "fmt"

func main() {
	fmt.Println(Maps([]int{10, 100}))
}

//Maps takes an array of int, double its elements, and return a new int array
func Maps(numbers []int) []int {
	var newNumbers []int
	for _, number := range numbers {
		newNumbers = append(newNumbers, number*2)
	}
	return newNumbers
}
func combat(health, damage float64) float64 {
	health -= damage
	if health <= 0.0 {
		return 0.0
	}
	return health
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
