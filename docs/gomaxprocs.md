# GOMAXPROCS Lever

Controls the **number of OS threads** that will host so-called `work queues`

```go
runtime.GOMAXPROCS(runtime.NumCPU())
```

- Prior to Go 1.5, `GOMAXPROCS` was always set to **1**
- Now automatically set to the **number of logical CPUs** on the host machine
