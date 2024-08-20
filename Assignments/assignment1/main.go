package main

import (
	"fmt"
	"strings"
)

// Create a program that has multiple string variable and displays the string on one line.
func main() {
	fmt.Println(concatenateStrings("one", "two", "three"))
}

func concatenateStrings(stringList ...string) string {
	var completeString string
	for _, stringText := range stringList {
		completeString += stringText + " "
	}
	return strings.TrimSpace(completeString)
}
