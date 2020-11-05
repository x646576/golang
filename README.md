# golang

- [golang.org](https://golang.org/)
  - [doc](https://golang.org/doc/)
  - [release](https://golang.org/dl/)
    - [install](https://golang.org/doc/install)
  - [packages](https://golang.org/pkg/)
- [gvm](https://github.com/moovweb/gvm): Go Version Manager

- Tutorial
  - [Get started with Go](https://golang.org/doc/tutorial/getting-started)
  - [Create a Go module](https://golang.org/doc/tutorial/create-module)

---

## Install

1. Install [gvm](https://github.com/moovweb/gvm)
1. Releases: `gvm install`
1. Install go: `gvm install goX.Y.Z --binary`
1. Go versions: `gvm list`
1. Select: `gvm use goX.Y.Z`
1. Env vars: `go env`

---

## Get started

### Create a project directory in $GOROOT/src/project-name

```bash
mkdir $GOROOT/src/hello
cd $GOROOT/src/hello
```

### Run a script

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

### Compile a script

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

### Install a executable

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
