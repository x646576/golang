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

code: [](../examples/patterns/.go)

## Error handling

code: [](../examples/patterns/.go)

## Pipelines

code: [](../examples/patterns/.go)

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
