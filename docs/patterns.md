# Patterns

## Confinement

code: [confinement](../src/patterns/confinement.go)

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

code: [for-select](../src/patterns/for-select.go)

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

code: [goroutine leaks](../src/patterns/goroutine-leaks.go), [goroutine leaks 2](../src/patterns/goroutine-leaks2.go)

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

code: [or-channel](../src/patterns/or-channel.go)

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

code: [error handling](../src/patterns/error-handling.go)

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

code: [pipelines](../src/patterns/pipelines.go), [generators](../src/patterns/pipelines-generators.go), [benchmark](../src/patterns/pipeline-benchmark_test.go)

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

- Fan-out: the process of starting multiple goroutines to handle input from the pipeline
- Fan-in: the process of combining multiple results into one channel

Consider fanning out:

- It doesn’t rely on values that the stage had calculated before.
- It takes a long time to run.

### Prime finder

code: [prime finder](../src/patterns/primefinder.go)

```bash
Primes:
        24941317
        36122539
        6410693
        10128161
        25511527
        2107939
        14004383
        7190363
        45931967
        2393161
Search took: 23.802461435s
```

### Fan-out

code: [fan-out prime finder](../src/patterns/fanout-primefinder.go)

```go
numFinders := runtime.NumCPU()
finders := make([]<-chan interface{}, numFinders)
for i := 0; i < numFinders; i++ {
  finders[i] = primeFinder(done, randIntStream)
}

fanIn := func(
  done <-chan interface{},
  channels ...<-chan interface{},
) <-chan interface{} {
  var wg sync.WaitGroup
  multiplexedStream := make(chan interface{})

  multiplex := func(c <-chan interface{}) {
    defer wg.Done()
    for i := range c {
      select {
      case <-done:
        return
      case multiplexedStream <- i:
      }
    }
  }

  wg.Add(len(channels)) // Select from all the channels
  for _, c := range channels {
    go multiplex(c)
  }

  go func() {
    wg.Wait() // Wait for all the reads to complete
    close(multiplexedStream)
  }()

  return multiplexedStream
}

for prime := range take(done, fanIn(done, finders...), 10) {
  fmt.Printf("\t%d\n", prime)
}
```

```bash
Spinning up 8 prime finders.
Primes:
        6410693
        24941317
        10128161
        36122539
        25511527
        2107939
        14004383
        7190363
        2393161
        45931967
Search took: 5.498295404s
```

## or-done-channel

code: [or-done](../src/patterns/or-done.go)

```go
orDone := func(done, c <-chan interface{}) <-chan interface{} {
  valStream := make(chan interface{})
  go func() {
    defer close(valStream)
    for {
      select {
      case <-done:
        return
      case v, ok := <-c:
        if ok == false {
          return
        }
        select {
        case valStream <- v:
        case <-done:
        }
      }
    }
  }()
  return valStream
}

for val := range orDone(done, myChan) {
  // do somthing with val
  fmt.Println(val)
}
```

## tee-channel

code: [tee-channel](../src/patterns/tee-channel.go)

```go
tee := func(
  done <-chan interface{},
  in <-chan interface{},
) (_, _ <-chan interface{}) {
  out1 := make(chan interface{})
  out2 := make(chan interface{})

  go func() {
    defer close(out1)
    defer close(out2)

    for val := range orDone(done, in) { // 1 2 1 2
      var out1, out2 = out1, out2
      for i := 0; i < 2; i++ { // 0 1 <- to ensure both are written
        select {
        case <-done:
        case out1 <- val:
          out1 = nil // block writing
        case out2 <- val:
          out2 = nil // block writing
        }
      }
    }
  }()
  return out1, out2
}
```

## bridge-channel

code: [](../src/patterns/.go)

## Queueing

code: [](../src/patterns/.go)

## Context package

code: [](../src/patterns/.go)
