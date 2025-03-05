package main

import "fmt"

func add(a, b int) int {
	fmt.Printf("%d and %d", a, b)

	return a + b
}

func multiply(a, b, c int) int {
	fmt.Printf("%d, %d and %d", a, b, c)

	return a * b * c
}

func main() {
	// Addition
	fmt.Printf("Addition of ")

	res := add(1, 2)
	fmt.Println(" is", res)

	// Multiplication
	fmt.Printf("Multiplication of ")

	res = multiply(1, 2, 3)
	fmt.Println(" is", res)
}
