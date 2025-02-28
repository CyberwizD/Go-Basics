# Variadic Functions
Usually, functions in Go accept only a fixed number of arguments. However, it is also possible to write variadic functions in Go.

A variadic function is a function that accepts a variable number of arguments.

If the type of the last parameter in a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.

```go
func find(a int, b ...int) {
    // ...
}
```

In the above function, parameter b is variadic and we can pass 0 or more arguments to b.

```go
find(5, 6)
find(5, 6, 7)
find(5)
```

## Caution
The variadic parameter must be the last parameter of the function.

The way variadic functions work is by converting the variable number of arguments to a slice of the type of the variadic parameter.

Sometimes you already have a slice and want to pass that to a variadic function. This can be achieved by passing the slice followed by ```...```. That will tell the compiler to use the slice as is inside the variadic function. The step described above where a slice is created will simply be omitted in this case.

It is important to note that ```...``` is not an actual rest and spread operator in Go. For example, the following code does not compile.

```go
func myFunc(a int, b int, c int) {
	// ...
}

func main() {
	nums := []int{1, 2, 3}
	myFunc(nums...)
}
```

This results in invalid use of ```...``` in call to myFunc because myFunc does not have a variadic parameter. ```...``` really only works in the specific scenarios explained above.
