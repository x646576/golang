package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  c := sync.NewCond(&sync.Mutex{})
  queue := make([]interface{}, 0, 10)

  removeFromQueue := func(delay time.Duration) {
    time.Sleep(delay)
    c.L.Lock() // critical section
    queue = queue[1:] // sub slice
    fmt.Println("Removed from queue")
    c.L.Unlock()
    c.Signal() // event
  }

  for i := 0; i < 10; i++ {
    c.L.Lock()
    for len(queue) == 2 { // check queue size
      c.Wait() // suspend main goroutine
    }
    fmt.Println("Adding to queue")
    queue = append(queue, struct{}{})
    go removeFromQueue(1 * time.Second)
    c.L.Unlock()
  }
}