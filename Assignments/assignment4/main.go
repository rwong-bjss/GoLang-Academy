package main

import "fmt"

// Create a program that initialises an array with the integer values 1 to 10:
// Display the array content in ascending sequential order 1 to 10.
// Display the array content in descending sequential order 10 to 1.
// Count even numbers and odd numbers in increasing and decreasing sequential order.
// Display the even and odd count sequences to screen.
func main() {
	fmt.Printf("ascending: %d\n", returnArrayAscending())
	fmt.Printf("descending: %d\n", returnArrayDescending())
	evenAsc, oddAsc, evenDesc, oddDesc := returnSeriesOfNumbers(returnArrayAscending(), returnArrayAscending(), returnArrayDescending(), returnArrayDescending())
	fmt.Printf("evenAsc: %d\n", evenAsc)
	fmt.Printf("oddAsc: %d\n", oddAsc)
	fmt.Printf("evenDesc: %d\n", evenDesc)
	fmt.Printf("evenAsc: %d\n", oddDesc)
}

func returnArrayAscending() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func returnArrayDescending() []int {
	ascendingList := returnArrayAscending()
	var descending []int
	for i := len(ascendingList) - 1; i >= 0; i-- {
		descending = append(descending, ascendingList[i])
	}
	return descending
}

func returnSeriesOfNumbers(evenAscending, oddAscending, evenDescending, oddDescending []int) ([]int, []int, []int, []int) {
	var evenAsc, oddAsc, evenDesc, oddDesc []int
	for i := 0; i < 10; i++ {
		if evenAscending[i]%2 == 0 {
			evenAsc = append(evenAsc, evenAscending[i])
		}
		if oddAscending[i]%2 != 0 {
			oddAsc = append(oddAsc, oddAscending[i])
		}
		if evenDescending[i]%2 == 0 {
			evenDesc = append(evenDesc, evenDescending[i])
		}
		if oddDescending[i]%2 != 0 {
			oddDesc = append(oddDesc, oddDescending[i])
		}
	}
	return evenAsc, oddAsc, evenDesc, oddDesc
}

//
//func deleteElement(slice []int, index int) []int {
//	return append(slice[:index], slice[index+1:]...)
//}
