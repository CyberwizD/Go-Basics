# Floating-Point Numbers in Go

A **floating-point number** is a number that can have zero or more digits after the decimal separator. Some examples of floating-point numbers are:

- `-2.4`
- `0.1`
- `3.14`
- `16.984025`
- `1024.0`

Different floating-point types can store different numbers of digits after the decimal point, which is referred to as their **precision**.

Go provides two floating-point types:

- **float32**: Uses 32 bits and has a precision of about 6 to 9 digits.
- **float64**: Uses 64 bits and has a precision of about 15 to 17 digits. This is the **default floating-point type** in Go.

For example, trying to store the value of PI in a `float32` variable will only store the first 6 to 9 digits, with the last digit being rounded.

### Precision of Float Types

- **float32**: ~6-9 digits of precision
- **float64**: ~15-17 digits of precision

### Type Defaults

By default, Go will use `float64` for floating-point numbers unless:

- The number is assigned to a variable with the type `float32`.
- The number is returned from a function with a return type `float32`.
- The number is passed as an argument to the `float32()` conversion function.

For example:

```go
var x float32 = 3.14   // x will be float32
y := 3.14              // y will be float64 by default

z := float32(3.14)     // z will be float32 using explicit conversion
```

In these cases, Go will handle the floating-point numbers accordingly based on their type or conversion.