// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

package main

import "fmt"

func main() {
	teste := [10]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}
	first := Points(teste)
	fmt.Println(teste) // 0/
	// fmt.Println(nextEven()) // 2
	// fmt.Println(nextEven()) // 4
}
func Points(games []string) string {
	game := games[0]
	return game
}
