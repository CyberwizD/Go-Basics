# Type Switches

A type switch can perform several type assertions in series. It has the same syntax as a type assertion (interfaceVariable.(concreteType)), but the concrete type is replaced with the keyword type.

Here is an example:
```go
var i interface{} = 12 // try: 12.3, true, int64(12), []int{}, map[string]int{}

switch v := i.(type) {
case int:
    fmt.Printf("the integer %d\n", v)
case string:
    fmt.Printf("the string %s\n", v)
default:
    fmt.Printf("type, %T, not handled explicitly: %#v\n", v, v)
}
```

In this example, the **switch** checks the underlying type of **i** and matches it to a case, executing the corresponding block. The **default** case handles any types that are not explicitly handled.
