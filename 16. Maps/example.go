package main

import (
	"fmt"
	"maps"
)

func main() {
	// Create a map
	m := make(map[string]int)

	// Assign key-value pair
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	// Delete a key from the map
	delete(m, "k2")
	fmt.Println("map:", m)

	// Clear the map
	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	// Check if 'n' map is equal to 'n2' map
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
