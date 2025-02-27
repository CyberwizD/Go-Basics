package main

import "fmt"

func main() {
	var number int
	result := "This number is "

	if number > 0 {
		result += "positive"
	} else if number < 0 {
		result += "negative"
	} else {
		result += "zero"
	}

	num := 7
	if v := 2 * num; v > 10 {
		fmt.Println(v)
	} else {
		fmt.Println(num)
	}
	// Output: 14

}
