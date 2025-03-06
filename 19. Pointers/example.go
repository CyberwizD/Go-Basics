package main

import "fmt"

// Function 'zeroval' takes an integer argument by value
// It does not modify the original value because it works on a copy of the variable.
func zeroval(ival int) {
	ival = 0 // This change is local to the function
}

// Function 'zeroptr' takes a pointer to an integer as an argument
// It modifies the original value because it operates on the memory address of the variable.
func zeroptr(iptr *int) {
	*iptr = 0 // Dereferences the pointer and assigns 0 to the original value
}

func main() {
	// Initialize a variable 'i' with the value of 1
	i := 1
	fmt.Println("initial:", i) // Output: initial: 1

	// Call 'zeroval' passing 'i' by value, meaning it won't affect the original variable 'i'
	zeroval(i)
	fmt.Println("zeroval:", i) // Output: zeroval: 1 (unchanged since 'zeroval' modifies a copy of 'i')

	// Call 'zeroptr' passing the memory address of 'i' (&i), allowing modification of the original variable
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // Output: zeroptr: 0 (the original variable 'i' is modified through the pointer)

	// Print the memory address of 'i'
	fmt.Println("pointer:", &i) // Output: pointer: <memory address of 'i'>
}
