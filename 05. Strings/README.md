# Strings in Go

A string in Go is an immutable sequence of bytes, which don't necessarily have to represent characters.

A string literal is defined between double quotes:

```go
const name = "Jane"
```

Strings can be concatenated via the + operator:

```go
"Jane" + " " + "Austen"
// => "Jane Austen"
```

Some special characters need to be escaped with a leading backslash, such as \t for a tab and \n for a new line in strings:

```go
"How is the weather today?\nIt's sunny"
// => How is the weather today?
// It's sunny
```

## Working with the strings Package
The strings package contains many useful functions to work on strings. 
