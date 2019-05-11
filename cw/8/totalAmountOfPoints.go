// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	teste := []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}
	first := Points(teste)
	fmt.Println(first)
}

// Points calculates the amount of poins the home team has
func Points(games []string) int {
	var scoreBoard int
	for index := 0; index < len(games); index++ {
		game := strings.Split(games[index], ":")
		homeScore, _ := strconv.Atoi(game[0])
		visitorScore, _ := strconv.Atoi(game[1])
		if homeScore > visitorScore {
			scoreBoard += 3
		} else if homeScore == visitorScore {
			scoreBoard++
		}
	}
	return scoreBoard
}
