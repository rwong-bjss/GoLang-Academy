package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create a program that lets the user input a first name, middle name and last name. Display the person's full name on one line.
func main() {
	pl("What is your first name?")
	firstName := readInput()
	pl("What is your middle name?")
	middleName := readInput()
	pl("What is your last name?")
	lastName := readInput()
	fullName := concatenateFullName(firstName, middleName, lastName)
	pl(fullName)
	f, s, l := sliceFullName(fullName)
	pl("Firstname: ", f)
	pl("SecondName: ", s)
	pl("LastName: ", l)
}

var reader = bufio.NewReader(os.Stdin)
var pl = fmt.Println

func concatenateFullName(first, middle, last string) string {
	return first + " " + middle + " " + last
}

func readInput() string {
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(name, "\n")
}

func sliceFullName(fullName string) (firstName, secondName, LastName string) {
	names := strings.Fields(fullName)
	if len(names) != 3 {
		log.Fatal("error name was not 3 names")
	}
	return names[0], names[1], names[2]
}
