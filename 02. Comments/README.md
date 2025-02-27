# Go Documentation Comments

In Go, comments play an important role in documenting code. They are used by the `godoc` command, which extracts these comments to create documentation for Go packages. A documentation comment should be a complete sentence that starts with the name of the thing being described and ends with a period.

Comments should precede packages as well as exported identifiers, for example exported functions, methods, package variables, constants, and structs.

---

## Package-Level Variable Example

A package-level variable might look like this:

```go
// TemperatureCelsius represents a certain temperature in degrees Celsius.
var TemperatureCelsius float64
```

## Package Comments

Package comments should be written directly before a package clause (package x) and begin with Package x ..., like this:

```go
// Package kelvin provides tools to convert
// temperatures to and from Kelvin.
package kelvin
```

## Function Comments
A function comment should be written directly before the function declaration. It should be a full sentence that starts with the function name. The comment should explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period.

For example:

```go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
    return 0
}
```
