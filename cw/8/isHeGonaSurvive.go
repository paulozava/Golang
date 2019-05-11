package main

import "fmt"

func main() {
	fmt.Println(Hero(12, 1))
}

func Hero(bullets, dragons int) bool {
	bulletsPerDragon := bullets / 2
	if bulletsPerDragon >= dragons {
		return true
	}
	return false
}
