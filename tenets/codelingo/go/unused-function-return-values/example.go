package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "incorrect string"

	strings.Trim(str, "in") // ISSUE
	fmt.Println(str)

	correct := strings.Trim(str, "in")
	fmt.Println(correct)
}
