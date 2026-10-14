package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	ID int
	Name string
	Age int
	Grades Grades
}

type Grades struct {
	Arabic float64
	English float64
	Math float64
	Science float64
}

func main() {
	class := []Student{}
	
	reader := bufio.NewReader(os.Stdin)

	outerLoop:
	for {
		var input string

		fmt.Println("Choose the number of what you want:")
		fmt.Println("0: Exit")
		fmt.Println("1: Add student")
		fmt.Println("2: Print students")
		fmt.Println("3: Search for a student by id")

		fmt.Scanln(&input)

		switch input {
		case "0":
			fmt.Println("Goodby!")
			break outerLoop;
		case "1":
			st := Student{}

			readStudent(reader, &st)

			class = append(class, st)
			fmt.Println("Student has been added successfully!")
		case "2":
			fmt.Println(class)
		case "3":
			fmt.Print("Enter student ID:")
			var id int
			fmt.Scan(&id)
			searchStudentsById(&class, id)
		}
	}
}

func readStudent(reader *bufio.Reader, st *Student){
	fmt.Print("Enter student name:")
	s, _ := reader.ReadString('\n')
	st.Name = s

	fmt.Print("Enter student ID:")
	fmt.Scan(&st.ID)

	fmt.Print("Enter student age:")
	fmt.Scan(&st.Age)

	fmt.Print("Enter student arabic grade:")
	fmt.Scan(&st.Grades.Arabic)

	fmt.Print("Enter student english grade:")
	fmt.Scan(&st.Grades.English)

	fmt.Print("Enter student math grade:")
	fmt.Scan(&st.Grades.Math)

	fmt.Print("Enter student science grade:")
	fmt.Scan(&st.Grades.Science)
}

func searchStudentsById(class *[]Student, id int) {
	if (len(*class) == 0) {
		fmt.Println("No students yet in the class")
		return
	}

	for _, st := range *class {
		if( st.ID == id) {
			fmt.Println(st)
		} else {
			fmt.Println("No student found with ID:", id)
		}
	}
}

// var st *Student