# Patterns

## Confinement

code: [confinement](../examples/patterns/confinement.go)

```go
printData := func(wg *sync.WaitGroup, data []byte) {
  defer wg.Done()

  var buff bytes.Buffer
  for _, b := range data {
    fmt.Fprintf(&buff, "%c", b)
  }
  fmt.Println(buff.String())
}

var wg sync.WaitGroup
wg.Add(1)
data := []byte("golang")
go printData(&wg, data[:3]) // lexical confinement
wg.Wait()
```

## for-select loop

code: [for-select](../examples/patterns/for-select.go)

```go
for {
  select {
    // channel job
  }
}
```

### Sending iteration variables out on a channel

```go
for _, s := range []string{"a", "b", "c"} {
  select {
  case <-done:
    return
  case stringStream <- s:
  }
}
```

### Looping infinitely waiting to be stopped

```go
for {
  select {
  case <-done:
    return
  default:
    // Do non-preemptable work
  }
  // Do non-preemptable work
}
```

## Preventing goroutine leaks

code: [goroutine leaks](../examples/patterns/goroutine-leaks.go), [goroutine leaks 2](../examples/patterns/goroutine-leaks2.go)

### example 1

```go
done := make(chan interface{})

go func() {
  time.Sleep(1 * time.Second)
  fmt.Println("Cancelling doWork goroutine...")
  close(done)
}()

doWork := func(done <-chan interface{}, strings <-chan string) {
  go func() {
    for {
      select { // Waiting...
      case s := <-strings: // nil channel
        fmt.Println(s)
      case <-done: // cancellation
        return
      }
    }
  }()
}

<-doWork(done, nil)
```

### example 2

```go
newRandStream := func(done <-chan interface{}) <-chan int {
  randStream := make(chan int)
  go func() {
    defer close(randStream)
    for {
      select {
      case <-done:
        return
      case randStream <- rand.Int():
      }
    }
  }()
  return randStream
}

done := make(chan interface{})
randStream := newRandStream(done)

for i := 1; i <= 3; i++ {
  <-randStream
}

close(done) // close goroutine
```

## or-channel

code: [or-channel](../examples/patterns/or-channel.go)

Combine one or more _done_ channels into a single _done_ channel that closes if any of its component channels close.

```go
<-or(
  sig(2*time.Hour),
  sig(5*time.Minute),
  sig(1*time.Second),
  sig(1*time.Hour),
  sig(1*time.Minute),
)

// done after 1.001742754s
```

## Error handling

code: [error handling](../examples/patterns/error-handling.go)

```go
type Result struct {
  Error    error
  Response *http.Response
}

var checkStatus func(done <-chan interface{}, urls ...string) <-chan Result

for result := range checkStatus(done, urls...) {
  if result.Error != nil {
    continue
  }
  // success
}
```

## Pipelines

code: [pipelines](../examples/patterns/pipelines.go), [generators](../examples/patterns/pipelines-generators.go), [benchmark](../examples/patterns/pipeline-benchmark_test.go)

```go
var generator func(done <-chan interface{}, integers ...int) <-chan int
var multiply func(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int
var add func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int

pipelines := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

for v := range pipelines {
}
```

### Benchmark

```bash
go test -bench=. pipeline-benchmark_test.go
```

```bash
BenchmarkGeneric-8        618374              1846 ns/op
BenchmarkTyped-8         1000000              1243 ns/op
PASS
ok      command-line-arguments  2.424s
```

## Fan-out, Fan-in

code: [](../examples/patterns/.go)

## or-done-channel

code: [](../examples/patterns/.go)

## tee-channel

code: [](../examples/patterns/.go)

## bridge-channel

code: [](../examples/patterns/.go)

## Queueing

code: [](../examples/patterns/.go)

## Context package

code: [](../examples/patterns/.go)
