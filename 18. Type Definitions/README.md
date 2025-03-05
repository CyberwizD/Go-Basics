# Struct Types in Go
To recap a **struct** is a sequence of named elements called **fields**, each field having a **name** and **type**. The name of a field must be unique within the struct. Structs can be compared with the class in the Object Oriented Programming paradigm.

You create a new struct by using the type and **struct** keywords, then explicitly define the name and type of the fields as shown in the example below.

```go
type StructName struct{
    field1 fieldType1
    field2 fieldType2
}
```

# Non-struct Types in Go

In Go, you've previously seen how to define struct types. However, it is also possible to define non-struct types, which you can use as an alias for a built-in type declaration. You can define receiver functions on these non-struct types in the same way as struct types to extend their behavior.

## Example: Defining a Non-struct Type

You can create a custom type based on a built-in type, such as `string`:

```go
type Name string

func SayHello(n Name) {
  fmt.Printf("Hello %s\n", n)
}

n := Name("Fred")
SayHello(n)
// Output: Hello Fred
```

## Example: Non-struct Type Composed of Arrays
You can also define non-struct types composed of arrays and maps. In the following example, we define a custom type Names, which is a slice of strings:

```go
type Names []string

func SayHello(n Names) {
  for _, name := range n {
    fmt.Printf("Hello %s\n", name)
  }
}

n := Names([]string{"Fred", "Bill"})
SayHello(n)
// Output:
// Hello Fred
// Hello Bill
```

This shows how non-struct types can be used to extend basic data types with methods that add functionality or enhance readability.
