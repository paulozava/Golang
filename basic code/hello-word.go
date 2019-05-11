package main

import "fmt"

var numbers = []int{1, 2, 3, 4, 5}

func add(a, b int) int {
	c := a + b
	return c
}

func main() {
	for index := 0; index < len(numbers); index++ {
		number := numbers[index]
		if number < 2 {
			added := add(12, number)
			fmt.Println(added)
		} else if number == 2 {
			fmt.Println("2 inside")
		} else {
			fmt.Println("Outer")
		}
	}
}
