package main

import (
	"strings"
)

func main() {
	RepeatString(10, "teste")
}
func RepeatString(repetitions int, word string) string {
	return strings.Repeat(word, repetitions)
}
