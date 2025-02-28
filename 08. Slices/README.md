# Go Slices

Slices in Go are similar to lists or arrays in other languages. They hold several elements of a specific type (or interface). 

## Overview
Slices in Go are based on arrays. Arrays have a fixed size, while slices are dynamically-sized, providing a flexible view of the elements of an array.

A slice is written like `[]T` with `T` being the type of the elements in the slice:

```go
var empty []int                 // an empty slice
withData := []int{0,1,2,3,4,5}  // a slice pre-filled with some data
```

You can get or set an element at a given zero-based index using square-bracket notation:

```go
withData[1] = 5
x := withData[1] // x is now 5
```

## Creating a New Slice
You can create a new slice from an existing one by getting a range of elements using the square-bracket notation. You can specify both a starting (inclusive) and an ending (exclusive) index. If you omit the starting index, it defaults to 0. If you omit the ending index, it defaults to the length of the slice.

```go
newSlice := withData[2:4]
// => []int{2,3}
newSlice := withData[:2]
// => []int{0,1}
newSlice := withData[2:]
// => []int{2,3,4,5}
newSlice := withData[:]
// => []int{0,1,2,3,4,5}
```

## Appending Elements to a Slice
You can add elements to a slice using the **append** function. Below, we append 4 and 2 to a slice:

```go 
a := []int{1, 3}
a = append(a, 4, 2)
// => []int{1,3,4,2}
```

**append** always returns a new slice. To append elements to an existing slice, it's common to reassign it back to the slice variable, as shown above.

append can also be used to merge two slices:

```go
nextSlice := []int{100,101,102}
newSlice  := append(withData, nextSlice...)
// => []int{0,1,2,3,4,5,100,101,102}
```

## Indexing in Slices
When working with slice indexes, always ensure the index exists by performing a check. Failing to do so will crash the application.

## Empty Slices
**nil** slices are the default empty slices. They function similarly to slices with no values. The **len** function works on **nil** slices, and items can be added without initializing them. To create a new slice, it's preferable to use **var s []int** (a **nil** slice) over **s := []int{}** (an empty, non-nil slice).

## Performance Consideration
When creating slices to be filled iteratively, there’s a simple optimization for performance if the final size of the slice is known. To minimize memory reallocation (which is expensive), specify the capacity (**cap**) of the slice using:

```go
s := make([]int, 0, cap)
```

Appending to this slice works as usual, but the memory space for cap items is allocated immediately, even though the slice length is zero.

In practice, **cap** is often set to the length of another slice:

```go
s := make([]int, 0, len(otherSlice))
```

## append is Not a Pure Function
The append function in Go is optimized for performance and does not make a copy of the input slice. This means that the original slice (the first parameter in append) might be modified during the operation.
