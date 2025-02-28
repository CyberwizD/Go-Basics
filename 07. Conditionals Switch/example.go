package main

func main() {
	operatingSystem := "windows"

	switch operatingSystem {
	case "windows":
		// Execute code if the operating system is Windows
	case "linux":
		// Execute code if the operating system is Linux
	case "macos":
		// Execute code if the operating system is macOS
	default:
		// Execute code if the operating system is none of the above
	}

	age := 21

	switch {
	case age > 20 && age < 30:
		// Execute code if age is between 20 and 30
	case age == 10:
		// Execute code if age is equal to 10
	default:
		// Execute code for every other case
	}
}
