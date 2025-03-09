# Stringer Interface and Example

The `Stringer` interface in Go is used to define the string format of values. It consists of a single method, `String()`, which returns a human-friendly string representation of the type.

```go
type Stringer interface {
    String() string
}
```

Types that want to implement this interface must define a `String()` method. Packages like `fmt` will look for this method to format and print values.

## Example: Distances
In this example, we work with geographical distances measured in different units. The types **DistanceUnit** and **Distance** are defined as follows:

```go
type DistanceUnit int

const (
	Kilometer    DistanceUnit = 0
	Mile         DistanceUnit = 1
)
 
type Distance struct {
	number float64
	unit   DistanceUnit
} 
```

Here, `Kilometer` and `Mile` are constants of type `DistanceUnit`.

## Default Formatting Behavior
Since these types do not implement the **Stringer** interface, calling `fmt.Sprint()` on them results in Go's default formatting.

```go
mileUnit := Mile
fmt.Sprint(mileUnit)
// => 1
```

The output is `1` because the underlying value of the **Mile** constant is `1` (as defined in the **const** block).

## For Distance:
```go
dist := Distance{number: 790.7, unit: Kilometer}
fmt.Sprint(dist)
// => {790.7 0}
// not a very useful output!
```

This is not very helpful as it prints the default structure format.

## Implementing the Stringer Interface
We can improve the output by implementing the **Stringer** interface for both **DistanceUnit** and **Distance**:

```go
func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
}
```

With these `String()` methods in place, `fmt` package functions will use them when formatting values.

```go
kmUnit := Kilometer
kmUnit.String()
// => km

mileUnit := Mile
mileUnit.String()
// => mi

dist := Distance{
	number: 790.7,
	unit: Kilometer,
}
dist.String()
// => 790.7 km
```

Now, the output is more meaningful and human-readable.

## Conclusion
By implementing the **Stringer** interface, we ensure that `fmt` functions output useful string representations for custom types. This is especially handy when dealing with types like **DistanceUnit** and **Distance** where the default formatting is not very informative.
