package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// fmt.Println(SumEvenFibonacci(1))
	teste := []int{1, 2, 2, -5, 5, -5}
	fmt.Println(multipleOfIndex(teste))
}

// multipleOfIndex Return a new array consisting of elements which are multiple of their own index in input array
func multipleOfIndex(ints []int) []int {
	var response []int
	for index := 1; index < len(ints); index++ {
		item := ints[index]
		if item%index == 0.0 {
			response = append(response, item)
		}
	}
	return response
}

//PartList return a coma moving string
func PartList(arr []string) string {
	var partial string
	for index := 1; index < len(arr); index++ {
		partial = partial + fmt.Sprintf("(%s, %s)", strings.Join(arr[:index], " "), strings.Join(arr[index:], " "))
	}
	return partial
}

// SumEvenFibonacci return a sum of all evens in a fibonacci's sequence
func SumEvenFibonacci(limit int) int {
	fibo := Fibonacci(limit)
	acc := 0
	for _, element := range fibo {
		if element%2 == 0 {
			acc += element
		}
	}
	return acc
}

//Fibonacci return a list with fibonacci sequence
func Fibonacci(limit int) []int {
	a, b := 1, 1
	fiboList := []int{1, 1}
	for b < limit {
		a, b = b, a+b
		fiboList = append(fiboList, b)
	}
	return fiboList
}

//EquableTriangle check if triangle is equable
func EquableTriangle(a, b, c int) bool {
	af, bf, cf := float64(a), float64(b), float64(c)
	perimeter := af + bf + cf
	semiPerimeter := perimeter / 2
	area := math.Sqrt(semiPerimeter * (semiPerimeter - af) * (semiPerimeter - bf) * (semiPerimeter - cf))
	return perimeter == area
}

//EvenOrOdd recive a number and return Even if its so, else return Odd
func EvenOrOdd(number int) string {
	if number%2 == 0.0 {
		return "Even"
	}
	return "Odd"
}

func century(year int) int {
	decade := int((year-1)/100) + 1
	return decade
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
