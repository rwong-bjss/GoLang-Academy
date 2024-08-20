package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create a program that allows the user to input a number. Check whether the number lies between 1 and 10.
func main() {
	pl("Input a number...")
	number := readInput()
	convertStringToInt(number)
	pl("This number is between 1 and 10 ", isNumberBetween1And10(convertStringToInt(number)))
}

var reader = bufio.NewReader(os.Stdin)
var pl = fmt.Println

func readInput() string {
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(name, "\n")
}

func isNumberBetween1And10(num int) bool {
	return 1 <= num && num <= 10
}

func convertStringToInt(value string) int {
	var i int
	if _, err := fmt.Sscan(value, &i); err == nil {
		return i
	} else {
		log.Fatal(err)
		//unreachable but required for the functions signature
		return 0
	}
}
