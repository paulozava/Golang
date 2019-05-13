package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("abcd"))
}

// RemoveChar  just a kata
func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
