package main

import "fmt"

// Define a struct 'person' with fields 'name' (string) and 'age' (int)
type person struct {
	name string
	age  int
}

// Function to create a new person instance, setting a default age of 42.
// Returns a pointer to a 'person' struct.
func newPerson(name string) *person {
	p := person{name: name} // Create a person instance with only the name set.
	p.age = 42              // Set the default age to 42.
	return &p               // Return a pointer to the person instance.
}

func main() {
	// Create a new 'person' instance with name "Bob" and age 20 and print it.
	fmt.Println(person{"Bob", 20})

	// Create a new 'person' instance using field names and print it.
	fmt.Println(person{name: "Alice", age: 30})

	// Create a new 'person' instance with name "Fred" and default age (zero value for int).
	fmt.Println(person{name: "Fred"})

	// Create a pointer to a 'person' instance with name "Ann" and age 40, and print it.
	fmt.Println(&person{name: "Ann", age: 40})

	// Use the 'newPerson' function to create a new person with name "Jon" and default age 42.
	fmt.Println(newPerson("Jon"))

	// Create a new 'person' instance with name "Sean" and age 50, then print the person's name.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Create a pointer to the person instance 's' and print the age using the pointer.
	sp := &s
	fmt.Println(sp.age)

	// Modify the age field of 'person' instance 's' using the pointer and print the new age.
	sp.age = 51
	fmt.Println(sp.age)

	// Define and create an anonymous struct for a 'dog' with fields 'name' and 'isGood', then print it.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex", // Set name to "Rex"
		true,  // Set isGood to true
	}
	fmt.Println(dog)
}
