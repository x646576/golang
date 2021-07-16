package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {
  var m sync.RWMutex
  var wg sync.WaitGroup

  const writer int = 1
  const reader int = 256

  wg.Add(writer + reader)
  beginTestTime := time.Now()
  go func() {
    defer wg.Done()
    for i := 5; i > 0; i-- {
      m.Lock()
      m.Unlock()
      time.Sleep(1)
    }
  }()
  for i := 0; i < reader; i++ {
    go func() {
      defer wg.Done()
      m.RLocker().Lock()
      defer m.RLocker().Unlock()
    }()
  }
  wg.Wait()

  fmt.Printf("\nReaders: %d\nTime: %v\n", reader, time.Since(beginTestTime))
}
