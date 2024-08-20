package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Create a program that: [Write File][Read File][I/O Package][I/O]
// Copies the following list of cities to a new file - "Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi".
// Reads a list of cities from the newly created file.
// Displays the list of cities in alphabetical order.
func main() {
	writeToFile(cities, "cities.txt")
	readFromFile("cities.txt")
}

var cities = []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

func writeToFile(contents []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()
	var concatString string
	for _, city := range contents {
		concatString += city + "\n"
	}
	_, err = file.WriteString(concatString)
	if err != nil {
		return
	}
}

func readFromFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	lines := strings.Split(string(data), "\n")
	fmt.Println(lines)
	sort.Strings(lines)
	fmt.Println(lines)
}
