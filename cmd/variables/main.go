package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Hello world!")

	// var lastName string // default is ""
	// var isMarried bool // default is false
	// var level int // default is 0

	// var name string = "Ahmed"
	// var age int = 30
	// var isStudent bool = true

	name := "Ahmed"
	age := 30
	isStudent := true
	grade := 3.5
	birthdate := time.Date(1996,6,20,0,0,0,0, time.UTC)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Is Student? %t\n", isStudent)
	fmt.Printf("Grade: %g\n", grade)
	fmt.Printf("Date: %s\n", birthdate.Format("2006-01-02"))
}