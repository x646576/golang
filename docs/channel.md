# Channel

Communicate information between goroutines.

- Direction: [Read and Write](../src/features/channel/readwrite.go), [Read and Write String Channel](../src/features/channel/readwriteString.go)
- Close: [Return second value](../src/features/channel/close.go), [for range](../src/features/channel/loop.go), [Wait empty channel](../src/features/channel/waitAndUnblockAll.go)
  - open: `true`
  - close: `false`
- [Buffer](../src/features/channel/buffer.go)
- [Channel Owner](../src/features/channel/owner.go)

## nil

Always initialize channel first.

### deadlock: read

```go
var dataStream chan interface{}
<-dataStream
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive (nil chan)]:
```

### write

```go
var dataStream chan interface{}
dataStream <- struct{}{}
```

```bash
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send (nil chan)]:
```

### panic: close of nil channel

```go
var dataStream chan interface{}
close(dataStream)
```
