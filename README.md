# golang

## Contents

- [Install](install.md): go version manager
- [Hello World](examples/helloworld/README.md): init, compile, run, test
- [Concurrency](docs/concurrency.md): race condition, deadlock conditions
  - [CSP](docs/csp.md): Communicating Sequential Processes
  - [goroutine](docs/goroutine.md), [sync](docs/sync.md), [channel](docs/channel.md), select
  - patterns
  - scaling
  - runtime
- Error, Race detection, pprof
- Use [Virtual Machine](docs/vm.md): vagrant

---

## Link

### [golang.org](https://golang.org/)

- Tutorial
  - [Get started with Go](https://golang.org/doc/tutorial/getting-started)
  - [Create a Go module](https://golang.org/doc/tutorial/create-module)
- [documentation](https://golang.org/doc/)
  - [Effective Go](https://golang.org/doc/effective_go)
- [modules](https://golang.org/ref/mod): Dependency management for Go
  - [Using Go Modules](https://blog.golang.org/using-go-modules)

### tools

- [gvm](https://github.com/moovweb/gvm): Go Version Manager

### book

- [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) by Katherine Cox-Buday
  - [Information](https://katherine.cox-buday.com/concurrency-in-go/)
  - [Source Code](https://github.com/kat-co/concurrency-in-go-src)

### blogs

- [You don't need virtualenv in Go](https://eli.thegreenplace.net/2020/you-dont-need-virtualenv-in-go/): `go.mod` file specifies the exact versions of dependency packages
- [The Free Lunch Is Over](http://www.gotw.ca/publications/concurrency-ddj.htm): We desperately need a higher-level programming model for concurrency than languages offer today
