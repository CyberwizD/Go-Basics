# Go Methods

A **method** is a function with a special receiver argument. The receiver appears in its own argument list between the `func` keyword and the name of the method.

```go
func (receiver type) MethodName(parameters) (returnTypes) {
    // do nothing
}
```

You can only define a method with a receiver whose type is defined in the same package as the method.

### Example: Value Receiver

```go
type Person struct {
    Name string
}

func (p Person) Greetings() string {
    return fmt.Sprintf("Welcome %s!", p.Name)
}

s := Person{Name: "Bronson"}
fmt.Println(s.Greetings())
// Output: Welcome Bronson!
```

## Types of Receivers: Value and Pointer Receivers
* **Value Receivers** operate on a copy of the value passed to it, meaning any modification done to the receiver inside the method is not visible to the caller.
* **Pointer Receivers** modify the value to which the receiver points. This is done by prefixing the type name with a `*`, for example, with the rect type, a pointer receiver would be declared as `*rect`. Modifications are visible to the caller of the method.

## Example: Pointer Receiver

```go
type rect struct {
    width, height int
}

func (r *rect) squareIt() {
    r.height = r.width
}

r := rect{width: 10, height: 20}
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 20

r.squareIt()
fmt.Printf("Width: %d, Height: %d\n", r.width, r.height)
// Output: Width: 10, Height: 10
```
