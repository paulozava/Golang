package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)

	fmt.Print("Try ")
	for scanner.Scan() {
		_, scannedError := strconv.ParseInt(scanner.Text(), 10, 10)
		if scannedError != nil {
			fmt.Println("bye bye")
			break
		}
		fmt.Print("Not ")
	}
}
