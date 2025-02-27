package main

import "fmt"

func main() {
	// Declaring variables

	var explicit int // Explicitly typed

	explicit = 1

	count := 1 // Assign initial value
	count = 2  // Update to new value

	// count = false // This throws a compiler error due to assigning a non `int` type

	fmt.Println(count, explicit)

	// Declaring constants

	const Age = 21          // Defines a numeric constant 'Age' with the value of 21
	const Name = "John Doe" // Defines a string constant 'Name' with the value of "John Doe"
}

// Hello is a public function.
func Hello(name string) string {
	return hi(name)
}

// hi is a private function.
func hi(name string) string {
	return "hi " + name
}
