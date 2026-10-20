package main

import "fmt"

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
	st := Student{}

	st.ID = 111222
	st.Name = "Ahmed Ouda"
	st.Age = 13
	st.Grades.Arabic = 90
	st.Grades.English = 80
	st.Grades.Math = 100
	st.Grades.Science = 85

	fmt.Println(st.ID)
	fmt.Println(st.Name)
	fmt.Println(st.Grades.Arabic)
}