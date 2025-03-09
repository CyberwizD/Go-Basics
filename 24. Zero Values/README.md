# Zero Values in Go

Go does not have a concept of `empty`, `null`, or `undefined` for variable values. When variables are declared without an explicit initial value, they default to the **zero value** of their respective type.

- **Primitive types** such as booleans, numeric types, and strings have the following zero values:
  - `boolean`: `false`
  - `numeric`: `0`
  - `string`: `""` (empty string)

For more complex types, the identifier `nil` is used as the zero value. Here is an overview of Go's zero values:

| Type      | Zero Value  |
|-----------|-------------|
| boolean   | `false`     |
| numeric   | `0`         |
| string    | `""`        |
| pointer   | `nil`       |
| function  | `nil`       |
| interface | `nil`       |
| slice     | `nil`       |
| channel   | `nil`       |
| map       | `nil`       |

## Zero Values for Structs

Structs are not directly listed in the table above because the zero value for a **struct** type depends on its fields. Each field of the struct is initialized to its respective zero value when the struct is declared without an initial value.

Example:

```go
type Person struct {
    Name string
    Age  int
}

var p Person
fmt.Println(p) // Output: {"" 0}
```

### In this case:
* The **Name** field defaults to an empty string `""`.
* The **Age** field defaults to `0`.

Understanding Go's zero values helps avoid unnecessary initialization and allows the use of default values across primitive and complex types.
