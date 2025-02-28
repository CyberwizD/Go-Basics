package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"

	strings.ToUpper(text)

	// `strings.ToLower` returns the string given as an argument with all its characters lowercased
	strings.ToLower("MaKEmeLoweRCase")
	// => "makemelowercase"

	// `strings.Repeat` returns a string with a substring given as an argument repeated many times
	strings.Repeat("Go", 3)
	// => "GoGoGo"

	// String Formatting

	food := "taco"
	message := fmt.Sprintf("Bring me a %s", food)
	fmt.Println(message)
	// Returns: Bring me a taco

	number := 4.3242
	formattedNumber := fmt.Sprintf("%.2f", number)
	fmt.Println(formattedNumber)
	// Returns: 4.32
}
