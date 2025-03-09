package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	// prints 9
	fmt.Println(aggregate(2, 3, 4, mul))
	// prints 24
}
