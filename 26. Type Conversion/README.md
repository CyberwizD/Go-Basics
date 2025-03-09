# Type Conversion in Go

Go requires explicit conversion between different types. Converting between types (also known as **type casting**) is done via a function with the name of the type to convert to. For example, to convert an `int` to a `float64`, you would need to do the following:

```go
var x int = 42 // x has type int
f := float64(x) // f has type float64 (i.e., 42.0)
```

## Converting between primitive types and strings
Go provides the strconv package for converting between primitive types (like int) and string.

### Example 1: Converting String to Integer
```go
import "strconv"

var intString string = "42"
var i, err = strconv.Atoi(intString) // Atoi converts string to int
```

### Example 2: Converting Integer to String
```go
import "strconv"

var number int = 12
var s string = strconv.Itoa(number) // Itoa converts int to string
```

## Note:
* `Atoi` stands for ASCII to integer.
* `Itoa` stands for integer to ASCII.

Always check for errors when converting, especially when converting from string to a numeric type.
