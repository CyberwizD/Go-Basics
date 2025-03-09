# The `error` Interface

Error handling in Go is not done via exceptions. Instead, errors are values of types that implement the built-in `error` interface. The `error` interface is very minimal, containing only one method `Error()` that returns the error message as a string.

```go
type error interface {
  Error() string
}
```

Whenever you define a function that could encounter an error, you need to include error as one of the return types. By convention, if the function has multiple return values, error is always the last one.

```go
func DoSomething() (int, error) {
  // ...
}
```

## Creating and Returning a Simple Error
You don't always need to implement the error interface yourself. To create a simple error, you can use the errors.`New()` function from the standard library errors package. This function takes the error message as a **string** and creates a **value** that contains your message and implements the error interface.

If the function returns an error, it is good practice to return the zero value for all other return parameters:

```go
func DoSomething() (SomeStruct, int, error) {
  // ...
  return SomeStruct{}, 0, errors.New("failed to calculate result")
}
```

## Caution
You should not assume that all functions return zero values for other return values if an error is present. It is best to assume that it is unsafe to use any of the other return values if an error occurred. The only exceptions are cases where the documentation explicitly states that other return values are meaningful in case of an error.

## Reusing Simple Errors
If you want to reuse an error message in multiple places, declare a variable for the error instead of using errors.New inline. By convention, such variables start with `Err` or `err` (depending on whether it is exported or not). These error variables are often called **sentinel errors**.

```go
import "errors"

var ErrNotFound = errors.New("resource was not found")

func DoSomething() error {
  // ...
  return ErrNotFound
}
```

To signal that there were no errors during function execution, return `nil` for the error:

```go
func Foo() (int, error) {
  return 10, nil
}
```

## Error Checking
When calling a function that returns an error, it is common to store the error value in a variable named `err`. Before using the actual result of the function, check whether an error occurred by comparing `err` to `nil`.

Handle error cases first to avoid nesting the **"happy path"** of your code. Use `==` and `!=` to compare the error against `nil`. There was an error if `err` **is not** `nil`.

```go
func processUserFile() error {
  file, err := os.Open("./users.csv")
  if err != nil {
    return err
  }

  // do something with file
}
```

Most of the time, the error will be returned up the function stack as shown above. Another way of handling the error could be to **log** it and continue with another operation. It is good practice to **either return or log the error**, **but never both**.

Since many functions in Go include an error as one of the return values, you will frequently encounter the `if err != nil` pattern.

## Custom Error Types
If you need your error to include more information than just a message string, you can create a **custom error type**. Any type that implements the error interface (i.e., has an **Error() string method**) can serve as an error in Go.

Usually, a **struct** is used to create a custom error type. By convention, custom error type names should end with `Error`. It is also recommended to set up the **Error()** string method with a **pointer receiver**. This ensures that when returning a custom error, you return a **pointer** to the **error**; otherwise, it will not implement the Error() method correctly.

```go
type MyCustomError struct {
  message string
  details string
}

func (e *MyCustomError) Error() string {
  return fmt.Sprintf("%s, details: %s", e.message, e.details)
}

func someFunction() error {
  // ...
  return &MyCustomError{
    message: "...",
    details: "...",
  }
}
```

In this example, the custom error `MyCustomError` includes both a **message** and additional **details**.
