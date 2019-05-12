package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AbbrevName("abraham lincoln"))
}
func AbbrevName(names string) string {
	namesArr := strings.Split(strings.ToUpper(names), " ")
	return fmt.Sprintf("%s.%s", namesArr[0][:1], namesArr[1][:1])
}
