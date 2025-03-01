package main

import (
	"fmt"
)

func main() {
	// Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For loop with condition
	j := 0

	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// Looping through an array
	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Looping through an array using range
	for i, v := range arr {
		fmt.Println(i, v)
	}

	// Looping through a map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
