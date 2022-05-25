# Goroutine

a function that is running concurrently.

Go program has at least one goroutine: _main goroutine_

When the goroutine has become blocked, goroutines are preemptable.

## closure

Wikipedia: [Closure](<https://en.wikipedia.org/wiki/Closure_(computer_programming)>)

a technique for implementing lexically scoped name binding in a language with first-class functions.

## Source Code

[src/runtime/proc.go](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/runtime/proc.go)

## Example

### goroutine

- [example/features/goroutine/hello.go](./../src/features/goroutine/hello.go)

Results:

```bash
Start
  3. Welcome
   4. hello (before Wait)
   4. welcome (after Wait)
1. Hello
    5. good day
    5. good day
    5. good day
     6. good day
     6. hello
     6. greetings
End
 2. Hi
```

### Memory of a goroutine

- [src/features/goroutine/memory.go](../src/features/goroutine/memory.go)

Result:

```bash
2.583kb
```

### Benchmark Context Switch

Use Linux or [Virtual Machine](vm.md)

### Install Packages

```bash
sudo apt install -y perf-tools-unstable linux-tools-5.4.0-65-generic linux-cloud-tools-5.4.0-65-generic linux-tools-generic linux-cloud-tools-generic
```

### Benchmark thread

Measures the time it takes to send and receive a message on a thread:

```bash
taskset -c 0 perf bench sched pipe -T
```

```bash
# Running 'sched/pipe' benchmark:
# Executed 1000000 pipe operations between two threads

     Total time: 5.764 [sec]

       5.764207 usecs/op
         173484 ops/sec
```

**Context Switch**: Divide it by two

- `5.764207 usecs/op` / 2
- `2.8821035 usecs/op`
- `2.8821035 μs per context switch`

### Benchmark goroutine

On macOS:

```bash
go test -bench=. -cpu=1 benchmark_test.go

goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkContextSwitch   6457441               175.2 ns/op
PASS
ok      command-line-arguments  1.328s
```

On Linux VM:

```bash
cd /src/features/goroutine
go test -bench=. -cpu=1 benchmark_test.go

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkContextSwitch   5725357               178.0 ns/op
PASS
ok      command-line-arguments  1.239s
```

Context Switch:

- `175.2 ns/op` ~ `178.0 ns/op`
- `0.175 μs per context switch`
