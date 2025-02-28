package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			return
		}
	}

	fmt.Println(num, "not found in ", nums)

	list := []int{1, 2, 3}

	find(1, list...) // "find" defined as shown above
	// Output:
	// type of nums is []int
	// 1 found at index 0 in [1 2 3]
}

func main() {
	// Here is an example of an implementation of a variadic function.

	find(89, 90, 91, 95)
	// =>
	// type of nums is []int
	// 89 not found in  [90 91 95]

	find(45, 56, 67, 45, 90, 109)
	// =>
	// type of nums is []int
	// 45 found at index 2 in [56 67 45 90 109]

	find(87)
	// =>
	// type of nums is []int
	// 87 not found in  []
}

/* In line find(89, 90, 91, 95) of the program above,
the variable number of arguments to the find function are 90, 91 and 95.
The find function expects a variadic int parameter after num.
Hence these three arguments will be converted by the compiler to a slice of type int []int{90, 91, 95}
and then it will be passed to the find function as nums. */
