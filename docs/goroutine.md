# Goroutine

a function that is running concurrently.

Go program has at least one goroutine: _main goroutine_

When the goroutine has become blocked, goroutines are preemptable.

## Source Code

[src/runtime/proc.go](https://cs.opensource.google/go/go/+/refs/tags/go1.16.6:src/runtime/proc.go)

## Example

- [example/features/goroutine](./../examples/features/goroutine/hello.go)

Results:

```bash
Start
Welcome
..............................End
```

```bash
Start
Welcome
..............................End
Hello
```

```bash
Start
Welcome
....Hi
......................Hello
....End
```
