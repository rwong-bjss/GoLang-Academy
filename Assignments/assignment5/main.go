package main

import (
	"fmt"
	"slices"
)

// Create a program that accepts and sums nine numbers. [Methods][Arrays][Slices][For loops]
// Three single digit numbers from one method.
// Three double digit numbers from a second method.
// Three triple digit numbers from a third method.
// Finally sum all methods into a final sum in the main program.
func main() {
	singleDigits := acceptSingleDigits()
	doubleDigits := acceptDoubleDigits()
	tripleDigits := acceptTripleDigits()
	fmt.Printf("single digits%d\n", singleDigits)
	allOfThem := slices.Concat(singleDigits, doubleDigits, tripleDigits)
	fmt.Printf("sum digits%d\n", sumAllNumbers(allOfThem))
}

// Todo fix this so it's better for being tested
func acceptSingleDigits() []int {
	numbers := make([]int, 3)
	fmt.Println("Enter three single-digit numbers:")
	for i := 0; i < 3; i++ {
		for {
			fmt.Printf("Number %d: ", i+1)
			fmt.Scan(&numbers[i])
			if numbers[i] >= 0 && numbers[i] <= 9 {
				break
			}
			fmt.Println("Please enter a single-digit number (0-9).")
		}
	}
	return numbers
}

func acceptDoubleDigits() []int {
	numbers := make([]int, 3)
	fmt.Println("Enter three double-digit numbers:")
	for i := 0; i < 3; i++ {
		for {
			fmt.Printf("Number %d: ", i+1)
			fmt.Scan(&numbers[i])
			if numbers[i] >= 10 && numbers[i] <= 99 {
				break
			}
			fmt.Println("Please enter a double-digit number (10-99).")
		}
	}
	return numbers
}
func acceptTripleDigits() []int {
	numbers := make([]int, 3)
	fmt.Println("Enter three triple-digit numbers:")
	for i := 0; i < 3; i++ {
		for {
			fmt.Printf("Number %d: ", i+1)
			fmt.Scan(&numbers[i])
			if numbers[i] >= 100 && numbers[i] <= 999 {
				break
			}
			fmt.Println("Please enter a triple-digit number (100-999).")
		}
	}
	return numbers
}
func sumAllNumbers(list []int) int {
	var sum int
	for _, num := range list {
		sum += num
	}
	return sum
}
