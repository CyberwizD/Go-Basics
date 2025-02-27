# Go Numbers

Go contains basic numeric types that can represent sets of either integer or floating-point values. This document covers two number types in Go:

---

## Number Types

### 1. `int`
- Examples: `0`, `255`, `2147483647`
- A signed integer that is at least 32 bits in size. This gives a value range of `-2147483648` through `2147483647`.
- On most modern 64-bit computers, an `int` will be 64 bits in size, with a value range of `-9223372036854775808` through `9223372036854775807`.

### 2. `float64`
- Examples: `0.0`, `3.14`
- Contains the set of all 64-bit floating-point numbers.

---

## Arithmetic Operators

Go supports the standard set of arithmetic operators:

- `+`: Addition
- `-`: Subtraction
- `*`: Multiplication
- `/`: Division
- `%`: Remainder (not modulo)

---

## Type Conversion in Go

In Go, assignment of a value between different types requires **explicit conversion**. For example, to convert an `int` to a `float64`.
