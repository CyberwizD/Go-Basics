# Maps in Go

In Go, a `map` is a built-in data type that maps keys to values. In other programming languages, you might be familiar with the concept of a map as a dictionary, hash table, key/value store, or an associative array.

Syntactically, a map looks like this:

```go
map[KeyType]ElementType
```

It is also important to know that each key is unique, meaning that assigning the same key twice will overwrite the value of the corresponding key.

## Creating a Map
To create a map, you can use either of the following:

```go
// With map literal
foo := map[string]int{}
```

or

```go
// With make function
foo := make(map[string]int)
```

## Map Operations
Here are some operations that you can do with a map:

### Add or Update a Value in a Map
```go
// Add a value in a map with the `=` operator
foo["bar"] = 42

// Here we update the element of `bar`
foo["bar"] = 73
```

## Retrieve a Map Value
```go
baz := foo["bar"]
```

## Delete an Item from a Map
```go
delete(foo, "bar")
```

## Checking for Existence of a Key
If you try to retrieve the value for a key that does not exist in the map, it will return the zero value of the value type. This can be confusing, especially if the default value of your ElementType (for example, 0 for an int) is a valid value. To check whether a key exists in your map, you can use:

```go
value, exists := foo["baz"]

// If the key "baz" does not exist,
// value: 0; exists: false
```
