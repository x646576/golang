# Select

```go
var c1, c2 <-chan interface{}
var c3 chan<- interface{}

select {
case <- c1:
  // job
case <- c2:
  // job
case c3<- struct{}{}:
  // job
}
```

## First

- [simple](../src/features/select/simple.go)

```go
select {
case <-c:
  fmt.Printf("Unblocked %v later.\n", time.Since(start))
}
```

```bash
Blocking on read...
Unblocked 5.002035092s later.
```

## What happens when multiple channels have something to read?

- [random](../src/features/select/random.go)

```go
select {
case <-c1:
  c1Count++
case <-c2:
  c2Count++
}
```

```bash
c1Count: 535
c2Count: 466
```

## What if there are never any channels that become ready?

- [time.After](../src/features/select/time-after.go)
- [default](../src/features/select/default.go)

```go
select {
case <-c1: // nil
case <-c2: // nil
case <-time.After(1 * time.Second):
  // job
// or
default:
  // job
}
```

## What if we want to do something but no channels are currently ready?

- [for & break](../src/features/select/for-break.go)

```go
loop:
	for {
		select {
		case <-done: // wait 5s
			break loop
		default:
		}

		workCounter++
		time.Sleep(1 * time.Second)
	}
```

## Block Forever

```go
select {}
```
