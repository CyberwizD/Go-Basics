package main

import "strings"

func main() {
	text := "Hello, World!"

	strings.ToUpper(text)

	// `strings.ToLower` returns the string given as an argument with all its characters lowercased
	strings.ToLower("MaKEmeLoweRCase")
	// => "makemelowercase"

	// `strings.Repeat` returns a string with a substring given as an argument repeated many times
	strings.Repeat("Go", 3)
	// => "GoGoGo"
}
