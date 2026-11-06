package main

import (
	"fmt"
	"strconv"
)

func main(){
	// 1. Example of bool with default value false
		
	// Declare a boolean variable without an explicit value
	// var isActive bool

	// fmt.Printf("Default value of isActive: %v (Type: %T)\n",isActive, isActive) // Default value of isActive: false (Type: bool)

	// -----------------------------------------
	// 2. Example of int and float with default value 0
	
	// Declare numeric variables without initial values
	// var count int
	// var price float64

	// fmt.Printf("Default value of count (Type: %T): %v\n", count, count) // Default value of count (Type: int): 0
	// fmt.Printf("Default value of price (Type: %T): %v\n", price, price) // Default value of price (Type: float64): 0

	// -----------------------------------------
	// 3. Example of Type Conversion between int and float
	
	// Convert from int to float64
	// var i int = 42
	// var f float64 = float64(i)

	// Example with float to int conversion WITH rounding
	// var f2 float64 = 99.99
	// var i2 int = int(f2)

	// math.Round() rounds to the nearest integer (returns a float64)
	// Then we convert that result to an int
	// var roundedInt int = int(math.Round(f2))

	// fmt.Printf("int (%d) converted to float64: %f\n", i, f) // int (42) converted to float64: 42.000000
	// fmt.Printf("float64 (%f) converted to int: %d (fraction lost)\n",f2,i2) // float64 (99.990000) converted to int: 99 (fraction lost)
	// fmt.Printf("Rounded and converted to int: %d (with math.Round())\n", roundedInt) // Rounded and converted to int: 100 (with math.Round())

	// -----------------------------------------
	// 4. Example of converting int and float to bool and vice versa

	// 1. Convert number to bool using conditions
	// score := 85
	// isPassed := score >= 50 // true

	// fmt.Printf("Score %d -> isPassed: %t\n", score, isPassed) // Score 85 -> isPassed: true

	// 2. Convert bool to number using if statement
	// isAdmin := true
	// var adminCode int
	// if isAdmin {
	// 	adminCode = 1
	// } else {
	// 	adminCode = 0
	// }

	// fmt.Printf("isAdmin %t -> adminCode: %d\n", isAdmin, adminCode) // isAdmin true -> adminCode: 1

	// -----------------------------------------
	// 5. Example of string with default empty string value

	// Declare a string without a value
	// var username string

	// fmt.Printf("Default value of username: %q (Length: %d)\n", username, len(username)) // Default value of username: "" (Length: 0)

	// -----------------------------------------
	// 6. Example of getting string length correctly and the difference with rune

	// text := "احمد"

	// Using standard len() (returns the number of bytes, not visual characters)
	// fmt.Printf("Byte length using len(): %d\n", len(text))

	// Using utf8 package to get the exact character count (Runes)
	// fmt.Printf("Actual rune count using utf8.RuneCountInString: %d\n", utf8.RuneCountInString(text))

	// -----------------------------------------
	// 7. Examples of manipulating strings using the strings package

	// s := "  Hello, Go Developers!  "

	// 1. Remove surrounding whitespace
	// trimmed := strings.TrimSpace(s)
	// fmt.Printf("strings.TrimSpace: %q\n", trimmed) // strings.TrimSpace: "Hello, Go Developers!"

	// 2. Convert text to uppercase
	// fmt.Printf("ToUpper: %s\n", strings.ToUpper(trimmed)) // ToUpper: HELLO, GO DEVELOPERS!

	// 3. Check if text contains a specific substring
	// fmt.Printf("Contains 'Go': %t\n", strings.Contains(trimmed, "Go")) // Contains 'Go': true
	
	// replaced := strings.ReplaceAll(trimmed, "Go", "Golang")
	// fmt.Printf("ReplaceAll: %s\n", replaced) // ReplaceAll: Hello, Golang Developers!

	// -----------------------------------------
	// 8. Example of converting string to byte/rune slice and vice versa

	// s := "Go Programming"
	
	// 1. Convert to a slice of bytes
	// byteSlice := []byte(s)
	// fmt.Printf("Bytes Slice: %v\n",byteSlice)
	// fmt.Printf("Back to string from bytes: %s\n", string(byteSlice))

	// 2. Convert to a slice of runes
	// runeSlice := []rune(s)
	// fmt.Printf("Runes Slice: %q\n",runeSlice)
	// fmt.Printf("Back to string from runes: %s\n", string(runeSlice))

	// -----------------------------------------
	// 9. Example of looping over a string using for

	// text := "أحمد"

	// fmt.Println("--- loop on range ---\n")
	// for index, char := range text {
	// 	fmt.Printf("Index: %d, Character: %c (Unicode: %U)\n", index, char, char)
	// }

	// fmt.Println("\n--- loop using traditional counter ---")
	// for i := 0; i < len(text); i++ {
	// 	b := text[i]
	// 	// %d prints the decimal value (0-255)
	// 	// 0x%x prints the hexadecimal value
	// 	// %c prints the raw byte as a character if it's printable ASCII
	// 	fmt.Printf("Byte index %d -> Decimal: %d | Hex: 0x%x\n", i, b, b)
	// }

	// -----------------------------------------
	// 10. Example of converting between string and bool

	// 1. From string to bool using strconv.ParseBool
	// strVal := "true"
	// b, err := strconv.ParseBool(strVal)
	// if err == nil {
	// 	fmt.Printf("Converted string %q to bool: %t (Type: %T)\n", strVal, b, b)
	// }

	// 2. From bool to string using strconv.FormatBool
	// isActive := false 
	// strResult := strconv.FormatBool(isActive)
	// fmt.Printf("Converted bool %t to string: %q (Type: %T)\n", isActive, strResult, strResult)

	// -----------------------------------------
	strInt := "2026"
	num, err := strconv.Atoi(strInt)
	if err == nil {
		fmt.Printf("String %q -> int: %d (Type: %T)\n", strInt, num, num)
	}

	strFloat := "3.14159"
	fNum, err := strconv.ParseFloat(strFloat, 64)
	if err == nil {
		fmt.Printf("String %q -> float64: %f (Type: %T)\n", strFloat,fNum,fNum)
	}

	age := 30
	ageStr := strconv.Itoa(age)
	fmt.Printf("Int %d -> string: %q (Type: %T)\n",age,ageStr,ageStr)

	price := 150.55
	priceStr := fmt.Sprintf("%.2f", price)
	fmt.Printf("Float %f -> string: %q (Type: %T)\n", price, priceStr, priceStr)
}