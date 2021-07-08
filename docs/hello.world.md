# Hello World

example code: [hello](../examples/hello/)

1. Create a project
1. Run a script
1. Compile a script
1. Install a executable

## Create a project directory in $GOROOT/src/project-name

```bash
mkdir $GOROOT/src/hello
cd $GOROOT/src/hello
```

## Run a script

```bash
vi hello.go
```

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

```bash
go run hello.go

Hello, World!
```

## Compile a script

```bash
go build hello.go
```

Result:

```bash
hello
├── [-rwxr-xr-x 2.0M]  hello
└── [-rw-r--r--   78]  hello.go
```

Run:

```bash
./hello

Hello, World!
```

## Install a executable

```bash
go install
```

Result:

```bash
hello
└── [-rw-r--r--   78]  hello.go
```

Location of the executable:

```bash
which hello
```

Run:

```bash
hello

Hello, World!
```
