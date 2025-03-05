package main

import "fmt"

func main() {
	// Create a slice
	nums := []int{2, 3, 4}

	// Initialize sum variable to be integer '0'
	sum := 0

	// Loop through slice 'nums'
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
