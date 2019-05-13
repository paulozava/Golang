package main

import (
	"strings"
)

func main() {
	RepeatString(10, "teste")
}

// RepeatString  just a kata
func RepeatString(repetitions int, word string) string {
	return strings.Repeat(word, repetitions)
}
