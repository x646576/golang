# Hello World

1. Init a go module
2. Write a code
3. Add a module
4. Run code
5. Compile and install a application
6. Test code

## Link

- Tutorial: [Get started with Go](https://golang.org/doc/tutorial/getting-started)
- Tutorial: [Create a Go module](https://golang.org/doc/tutorial/create-module)

---

## Init

```bash
helloworld/
├── greetings/
└── hello/
```

### greetings

```bash
cd greetings
go mod init example.com/greetings
```

### hello

```bash
cd hello
go mod init example.com/hello
```

### directory

```bash
helloworld/
├── greetings/
│   └── go.mod
└── hello/
    └── go.mod
```

---

## Code

### greetings.go

`greetings/greetings.go`

`Hello` returns a greeting for the named person.

```go
package greetings

import "fmt"

func Hello(name string) string {
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```

### hello.go

`hello/hello.go`

```go
package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    message := greetings.Hello("Obiwan")
    fmt.Println(message)
}
```

## Edit a module path

For now, because you haven't published the module yet,  
redirect `example.com/greetings` location.

```bash
cd hello
go mod edit -replace example.com/greetings=../greetings
```

### Add a requirement

```bash
cd hello
go mod tidy
```

### hello/go.mod

```bash
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

---

## Run

```bash
go run .
go run hello.go

Hi, Obiwan. Welcome!
```

---

## Compile a code

```bash
cd hello
go build hello.go
```

Result:

```bash
hello/
├── [-rw-r--r--  147]  go.mod
├── [-rwxr-xr-x 1.9M]  hello
└── [-rw-r--r--  145]  hello.go
```

Run:

```bash
./hello

Hi, Obiwan. Welcome!
```

---

## Install a application

### Install path

Discover the Go install path, where the go command will install the current package

```bash
go list -f '{{.Target}}'
```

### Compile and Install

```bash
go install
```

Result:

```bash
├── [-rw-r--r--  147]  go.mod
└── [-rw-r--r--  145]  hello.go
```

### Location of the executable

```bash
which hello
```

### Run the application

```bash
hello

Hi, Obiwan. Welcome!
```

---

## Test

Update code:

- [greetings/greetings.go](greetings/greetings.go)
- [greetings/greetings_test.go](greetings/greetings_test.go)

Run test:

```bash
cd greetings
go test -v
```

```bash
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)

PASS
ok      example.com/greetings   0.005s
```
