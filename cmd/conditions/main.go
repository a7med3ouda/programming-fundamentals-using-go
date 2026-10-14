package main

import "fmt"

// func main(){
// 	var age int

// 	fmt.Print("Enter your age: ")

// 	fmt.Scan(&age)

// 	if age >= 30 {
// 		fmt.Printf("Old Adult person with age %d\n", age)
// 	} else if age <30 && age >=18 {
// 		fmt.Printf("Young Adult person with age %d\n", age)
// 	} else {
// 		fmt.Printf("Child person with age %d\n", age)
// 	}
// }

func main(){
	var input string

	for {
		fmt.Print("Enter day of week:")

		fmt.Scan(&input)

		if(input == "exit") {
			fmt.Println("Goodbye!")
			break
		}

		switch input {
		case "fri", "sat":
			fmt.Println("Holly day")
		default:
			fmt.Println("Normal working day")
		}
	}
	// if (dayOfWeek == "fri" || dayOfWeek == "sat"){
	// 	fmt.Println("Holly day")
	// } else {
	// 	fmt.Println("Normal working day")
	// }
}