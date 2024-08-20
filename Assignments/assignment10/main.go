package main

import (
	"fmt"
	"time"
)

// Create a school register program that lists 10 pupils - full name, date of birth and age. [Structures][Arrays][Interfaces]
func main() {
	addStudent("rsw", time.Date(1997, 03, 24, 0, 0, 0, 0, time.UTC), 27)
	addStudent("dtudent 2", time.Date(1996, 03, 22, 0, 0, 0, 0, time.UTC), 28)
	addStudent("sally", time.Date(1997, 01, 24, 0, 0, 0, 0, time.UTC), 27)
	addStudent("bob bob bob", time.Date(1997, 12, 20, 0, 0, 0, 0, time.UTC), 27)
	addStudent("Mad Person", time.Date(1996, 02, 24, 0, 0, 0, 0, time.UTC), 28)
	students := listStudents()

	for _, student := range students {
		fmt.Printf("Name: %s, Date of Birth: %s, Age: %d\n",
			student.fullName,
			student.dateOfBirth.Format("2006-01-02"),
			student.age)
	}
}

var roll []person

type person struct {
	number      int
	fullName    string
	dateOfBirth time.Time
	age         int
}

type register interface {
	addStudent() person
	listStudents() []person
}

func listStudents() []person {
	return roll
}

func addStudent(fullName string, dob time.Time, age int) {
	p := person{len(roll) + 1, fullName, dob, age}
	roll = append(roll, p)
}
