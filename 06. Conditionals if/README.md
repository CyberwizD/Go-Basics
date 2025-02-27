# Conditionals in Go

Conditionals in Go are similar to conditionals in other languages. The underlying type of any conditional operation is the `bool` type, which can have the value of `true` or `false`. Conditionals are often used as flow control mechanisms to check for various conditions.

## If Statements

For checking a particular case, an `if` statement can be used, which executes its code if the underlying condition is `true`, like this:

```go
var value string

if value == "val" {
    return "was val"
}
```

## Else If and Else
In scenarios involving more than one case, many if statements can be chained together using the else if and else statements.

# Initialization in If Statements
If statements can also include a short initialization statement that can be used to initialize one or more variables for the if statement.

Note: Any variables created in the initialization statement go out of scope after the end of the if statement.