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
- [How to run](#how-to-run)

## Variables and Mutability

> Lesson file: [`src/classes/basic/1_variables_and_mutability.go`](src/classes/basic/1_variables_and_mutability.go)

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
