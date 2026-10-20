package main

import "fmt"

// func main() {

// 	fmt.Print("Enter first name:")

// 	var input string

// 	fmt.Scan(&input)

// 	fmt.Print("Enter last name:")

// 	var input2 string

// 	fmt.Scan(&input2)

// 	fmt.Printf("Welcome Mr/Ms %s %s to our application\n", input, input2)

// 	fmt.Printf("Input 1 address is %p\n", &input)
// 	fmt.Printf("Input 2 address is %p\n", &input2)
// }


func main() {
	var name *string = new(string)

	*name =  "Ahmed"

	printName(name)

	fmt.Printf("main: Address of variable name is %p\n", name)

	fmt.Println("main:", *name)
}

func printName(name *string) {
	*name = "Ouda"

	fmt.Println("printName:", *name)

	fmt.Printf("printName: Address of variable name is %p\n", name)
}