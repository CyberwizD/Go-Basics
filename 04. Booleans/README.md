# Booleans in Go

In Go, booleans are represented by the `bool` type. A `bool` is either `true` or `false`.

## Boolean Operators

Go supports three boolean operators: `!` (NOT), `&&` (AND), and `||` (OR).

### Examples:

```go
true || false  // => true
true && false  // => false
!true          // => false
```

## Operator Precedence
The three boolean operators have different precedence levels. They are evaluated in this order:

! (NOT) - highest precedence
&& (AND) - second precedence
|| (OR) - lowest precedence

### Examples:

```go
!true && false   // => false
!(true && false) // => true
```

## Parentheses for Custom Order
To force a different ordering, you can use parentheses (). Parentheses have a higher precedence than the boolean operators.

```go
!(true && false) // => true
```

By using parentheses, you can explicitly control how expressions are evaluated.
