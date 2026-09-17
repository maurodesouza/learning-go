# learning-go

## Overview

This project is a hands-on guide to learning how to **read and write Go**. It is
organized as a sequence of incremental lessons ("classes"), each focused on a
single topic. Lessons live under `src/classes/` and are run one at a time from
`src/main.go`, so you can follow along, tweak the code, and run the tests to see
how the language behaves.

## Contents

- [Overview](#overview)
- [Variables and Mutability](#variables-and-mutability)
- [Loops](#loops)
- [Functions](#functions)
- [How to run](#how-to-run)

## Variables and Mutability

> Lesson file: [`src/classes/basic/variables_and_mutability.go`](src/classes/basic/variables_and_mutability.go)

A quick recap of the topic covered in lesson 1.

### Three ways to declare

```go
const name = "value" // immutable, cannot be reassigned
var age = 30         // mutable, type inferred from the value
city := "Lisbon"     // short declaration, only valid inside a function
```

### Zero values

Every variable has an initial value. If you don't set one, Go uses the **zero
value** for its type:

```go
var s string  // ""
var i int     // 0
var b bool    // false
var f float64 // 0.0
```

### Mutability

Variables declared with `var` or `:=` can be reassigned; `const` cannot.

```go
count := 1
count = 2 // ok

const max = 10
// max = 20 // compile error: cannot assign to constant
```

### Short declaration rules

`:=` only works inside a function. It can also reassign existing variables, as
long as **at least one new variable** is introduced on the left side:

```go
a := "test"
a, b := "modified", 2 // ok: b is new, so a is reassigned
```

[Back to top](#learning-go)

## Loops

> Lesson file: [`src/classes/basic/loops.go`](src/classes/basic/loops.go)

A quick recap of the topic covered in lesson 2.

Go has a single loop keyword, `for`, which covers every looping style.

### Condition ("while")

Runs while the condition is true — Go's version of a `while` loop:

```go
for condition {
}
```

### Classic for

The traditional three-part loop:

```go
for initialization; condition; update {
}
```

### Range

Iterates over a collection (slice, map, string, ...), yielding the index/key
and the value:

```go
for index, value := range collection {
}
```

### Infinite

Loops forever until a `break` (or `return`) is reached:

```go
for {
}
```

[Back to top](#learning-go)

## Functions

> Lesson file: [`src/classes/basic/funcs.go`](src/classes/basic/funcs.go)

A quick recap of the topic covered in lesson 3.

### Functions are values

Functions can be assigned to variables or invoked immediately (IIFE).
Consecutive params of the same type can share the type declaration:

```go
sum := func(a, b int) int {
	return a + b
}
sum(1, 2) // 3

// IIFE: declared and called at once
result := func(a, b int) int {
	return a + b
}(1, 2)
```

### Variadic params

`...T` accepts any number of arguments and receives them as a slice. Spread
a slice into the call with `slice...`:

```go
func sum(numbers ...int) int { ... }

sum(1, 2, 3, 4, 5)
sum(slice...)
```

### Multiple and named returns

A function can return several values. Naming the returns lets a bare
`return` yield them:

```go
func calc(numbers ...int) (int, int) {
	return sumTotal, multiplyTotal
}

func calc(numbers ...int) (sumTotal int, multiplyTotal int) {
	// ...
	return // naked return
}
```

### Closures and higher-order functions

Functions can take and return other functions. The returned function
captures variables from the enclosing scope:

```go
counter := func() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}()

counter() // 1
counter() // 2
```

### defer

`defer` schedules a call to run when the surrounding function returns —
after the return value is computed, before the caller resumes:

```go
defer first()
second() // prints "second", then deferred "first" runs on return
```

[Back to top](#learning-go)

## How to run

Run the current lesson:

```bash
go run src/main.go
```

Run the tests:

```bash
go test ./...
```

[Back to top](#learning-go)
