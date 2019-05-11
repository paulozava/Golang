package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) describe() {
	fmt.Printf("The person %v has %v years\n", p.name, p.age)
}

func main() {
	hick := &person{name: "Hick", age: 42, gender: "Male"}
	hick.describe()
}
