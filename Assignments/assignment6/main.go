package main

import (
	"fmt"
	"github.com/bearbin/go-age"
	_ "github.com/bearbin/go-age"
	"time"
)

// Create a program that calculates the age of a person given their date of birth.
// [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
// (Use the github.com/bearbin/go-age to aid in the creation of this app.
// Also review unit testing applied to the application age.go within the imported package.)
func main() {
	calculate := calculateAgeFromToday(1997, 03, 24)
	fmt.Println(calculate)
}

func calculateAgeFromToday(year, month, day int) int {
	birth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return age.Age(birth)
}

func calculateAgeFrom(year, month, day int, now time.Time) int {
	//resetting for now
	now = time.Now()
	birth := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return age.AgeAt(birth, now)
}
