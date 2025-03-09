# First-Class Functions in Go

In Go, functions are **first-class values**. This means that you can do with functions the same things you can do with all other values—assign functions to variables, **pass them as arguments to other functions**, or even **return functions from other functions**.

Below we create two functions, `engGreeting` and `espGreeting`, and assign them to the variable `greeting`:

```go
import "fmt"

func engGreeting(name string) string {
    return fmt.Sprintf("Hello %s, nice to meet you!", name)
}

func espGreeting(name string) string {
    return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := engGreeting // `greeting` is a variable of type `func(string) string`
fmt.Println(greeting("Alice")) // Output: Hello Alice, nice to meet you!

greeting = espGreeting
fmt.Println(greeting("Alice")) // Output: ¡Hola Alice, mucho gusto!
```

Function values provide an opportunity to parameterize functions not only with data but with behavior too. In the following example, we pass behavior to the dialog function via the `greetingFunc` parameter:

```go
func dialog(name string, greetingFunc func(string) string) {
    fmt.Println(greetingFunc(name))
    fmt.Println("I'm a dialog bot.")
}

func espGreeting(name string) string {
    return fmt.Sprintf("¡Hola %s, mucho gusto!", name)
}

greeting := espGreeting
dialog("Alice", greeting)
// Output:
// ¡Hola Alice, mucho gusto!
// I'm a dialog bot.
```

## Nil Function Values
The value of an uninitialized variable of function type is nil. Therefore, calling a nil function value causes a panic:

```go
var dutchGreeting func(string) string
dutchGreeting("Alice") // panic: call of nil function
```

Function values can be compared with `nil`, which is useful to avoid unnecessary program panics:

```go
var dutchGreeting func(string) string
if dutchGreeting != nil {
    dutchGreeting("Alice") // safe to call dutchGreeting
}
```

## Function Types
Using function values is possible thanks to **function types** in Go. A function type denotes the set of all functions **with the same** sequence of parameter types and result types. **User-defined types** can be declared on top of function types. For instance, the dialog function from the previous example can be updated as follows:

```go
type greetingFunc func(string) string

func dialog(name string, f greetingFunc) {
    fmt.Println(f(name))
    fmt.Println("I'm a dialog bot.")
}
```

## Anonymous Functions
Another powerful tool enabled by first-class functions is **anonymous functions**. Anonymous functions are defined at their point of use without a name following the func keyword. Such functions **have access to the variables of the enclosing function**.

For example:

```go
func fib() func() int {
    var n1, n2 int

    return func() int {
        if n1 == 0 && n2 == 0 {
            n1 = 1
        } else {
            n1, n2 = n2, n1 + n2
        }
        return n2
    }
}

next := fib()
for i := 0; i < N; i++ {
    fmt.Printf("F%d\t= %4d\n", i, next())
}
```

A call to `fib` declares the variables `n1` and `n2` and returns an **anonymous function** that changes the values of these variables each time the function is called. The **N-th** call of the anonymous function returns the **N-th Fibonacci number** starting from `0`.

The anonymous inner function has access to the local variables (`n1` and `n2`) of the enclosing function `fib`. This is a great way to have function values keep state between calls. We say that the anonymous function is a closure of the variables `n1` and `n2`. **Closures** are widely used in programming, and other languages also support them.
