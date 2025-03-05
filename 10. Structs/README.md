# Structs in Go

In Go, a struct is a sequence of named elements called **fields**, where each field has a name and type. The name of a field must be unique within the struct. Structs can be compared with classes in the Object-Oriented Programming paradigm.

## Defining a Struct
You create a new struct by using the `type` and `struct` keywords, and explicitly define the name and type of the fields. 

For example, here we define a `Shape` struct that holds the name of a shape and its size:

```go
type Shape struct {
    name string
    size int
}
```

### Field names in structs follow Go convention:

* Fields whose name starts with a lowercase letter are only visible to code in the same package.
* Fields whose name starts with an uppercase letter are visible in other packages.

## Creating Instances of a Struct
Once you have defined the struct, you need to create a new instance defining the fields using their field name in any order:

```go
s := Shape{
    name: "Square",
    size: 25,
}
```

## Reading and Modifying Fields
To read or modify instance fields, use the . notation:

```go
// Update the name and the size
s.name = "Circle"
s.size = 35
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name: Circle size: 35
```

## Creating Instances Without Field Names
You can create an instance of a struct without using the field names, as long as you define the fields in order:

```go
s := Shape{
	"Oval",
	20,
}
```

However, this syntax is considered brittle code since it can break when a field is added, especially when the new field is of a different type. For example, let's add an extra field to Shape:

```go 
type Shape struct {
    name        string
    description string // new field 'description' added
    size        int
}

s := Shape{
    "Circle",
    20,
}
// Output: cannot use 20 (type untyped int) as type string in field value
// Output: too few values in Shape{...}
```

## New Functions
Sometimes it can be useful to have functions that help us create struct instances. By convention, these functions are usually called New or have their names starting with New. These are just regular functions and are not constructors like in some other languages.

For example, a NewShape function creates a new instance of Shape and automatically sets a default value for the size:

```go 
func NewShape(name string) Shape {
    return Shape{
        name: name,
        size: 100, // default-value for size is 100
    }
}

NewShape("Triangle")
// => Shape { name: "Triangle", size: 100 }
```

## Advantages of Using New Functions:
* Validation of the given values
* Handling of default values
* Since **New** functions are often declared in the same package as the structs they initialize, they can initialize even private fields of the struct.

## The **new** Built-in
```go
s := new(Shape) // s will be of type *Shape (pointer to Shape)
fmt.Printf("name: %s size: %d\n", s.name, s.size)
// Output: name:  size: 0
```

In this example, new creates an instance of the struct Shape with all the values initialized to the zero value of their type and returns a pointer to it.

## Note
More often than not, you will not see new instances of structs created using the new built-in. Always prefer using a struct literal, a custom New function, or any other function provided by the package that can help you build a struct. Use the new built-in to create new instances of structs as a last resort.

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

## References
- [Go by Example: Structs](https://gobyexample.com/structs)
- [A Tour of Go: Structs](https://tour.golang.org/moretypes/2)
- [Structures in Go (structs)](https://golangbot.com/structs/)
