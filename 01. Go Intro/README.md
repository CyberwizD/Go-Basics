# Go Basics

The **Basics** concept introduces Go as a statically typed, compiled programming language. The language is often referred to as *Golang* because of its previous domain name, golang.org, but the proper name is **Go**.

The Basics concept introduces three major language features:
- **Packages**
- **Functions**
- **Variables**

---

## Packages

Go applications are organized in **packages**. A package is a collection of source files located in the same directory. All source files in a directory must share the same package name. It is conventional for the package name to be the last directory in the import path.

For example, the files in the `"math/rand"` package begin with the statement `package rand`. When a package is imported, only entities (functions, types, variables, constants) whose names start with a capital letter can be accessed.

The recommended style of naming in Go is that identifiers will be named using `camelCase`, except for those meant to be accessible across packages which should be `PascalCase`.

```go
package lasagna
```

## Variables
Go is statically-typed, meaning all variables must have a defined type at compile-time.

Variables can be defined by explicitly specifying a type

## Constants 
Constans hold a piece of data just like variables, but their value cannot change during the execution of the program.

Constants are defined using the const keyword and can be numbers, characters, strings, or booleans:

## Functions
Go functions accept zero or more parameters. Parameters must be explicitly typed; there is no type inference.

Values are returned from functions using the return keyword. A function is invoked by specifying the function name and passing arguments for each of the function's parameters.

Note that Go supports two types of comments:

Single line comments are preceded by //
Multiline comments are inserted between /* and */
