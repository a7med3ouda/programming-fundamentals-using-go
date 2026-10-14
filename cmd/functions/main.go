package main

import "fmt"

// func main() {
// 	var a float64
// 	var b float64

// 	fmt.Print("Enter the first number:")

// 	fmt.Scan(&a)

// 	fmt.Print("Enter the second number:")

// 	fmt.Scan(&b)

// 	fmt.Println("Choose the operator:")
// 	fmt.Println("1: Plus +")
// 	fmt.Println("2: Minus -")

// 	var o string

// 	fmt.Scan(&o)

// 	switch o {
// 	case "1", "+":
// 		fmt.Println("The value is:", a + b)
// 	case "2", "-":
// 		fmt.Println("The value is:", a - b)
// 	}
// }

func main() {
	a := scanNumber("first")
	b := scanNumber("second")

	o := scanOperator()

	result := calculate(a,b,o)
	
	printTheResult(result)
}

func scanNumber(str string) float64 {
	var num float64

	fmt.Printf("Enter the %s number:", str)

	fmt.Scan(&num)

	return num
}

func scanOperator() string {
	var o string

	fmt.Println("Choose the operator:")
	fmt.Println("1: Plus +")
	fmt.Println("2: Minus -")

	fmt.Scan(&o)

	return o
}

func calculate(a float64, b float64, o string) float64 {
	switch o {
	case "1", "+":
		return a + b
	case "2", "-":
		return a - b
	default:
		panic("Wrong operator")
	}
}

func printTheResult(result float64){
	fmt.Println("The value is:", result)
}