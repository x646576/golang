# [Effective Go](https://go.dev/doc/effective_go)

## Keywords

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Table of Contents

- [Specification](https://golang.org/ref/spec)
  - [Keywords](https://golang.org/ref/spec#Keywords)

<details>
    <summary>Table of Contents</summary>

- Introduction
  - Examples
- Formatting
- Commentary
- Names
  - Package names
  - Getters
  - Interface names
  - MixedCaps
- Semicolons
- Control structures
  - If
  - Redeclaration and reassignment
  - For
  - Switch
  - Type switch
- Functions
  - Multiple return values
  - Named result parameters
  - Defer
- Data
  - Allocation with new
  - Constructors and composite literals
  - Allocation with make
  - Arrays
  - Slices
  - Two-dimensional slices
  - Maps
  - Printing
  - Append
- Initialization
  - Constants
  - Variables
  - The init function
- Methods
  - Pointers vs. Values
- Interfaces and other types
  - Interfaces
  - Conversions
  - Interface conversions and type assertions
  - Generality
  - Interfaces and methods
- The blank identifier
  - The blank identifier in multiple assignment
  - Unused imports and variables
  - Import for side effect
  - Interface checks
- Embedding
- Concurrency
  - Share by communicating
  - Goroutines
  - Channels
  - Channels of channels
  - Parallelization
  - A leaky buffer
- Errors
  - Panic
  - Recover
- A web server

</details>

---

### Formatting

```bash
gofmt -w a.go
```

- Indentation: tabs
- Line length: no limit

### Commentary
