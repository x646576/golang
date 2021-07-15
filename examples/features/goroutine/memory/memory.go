package main

import (
  "fmt"
  "sync"
  "runtime"
)

func main() {

  memConsumed := func () uint64 {
    runtime.GC()
    var s runtime.MemStats
    runtime.ReadMemStats(&s)
    return s.Sys
  }

  var c <-chan interface{}
  var wg sync.WaitGroup
  noop := func() { wg.Done(); <-c } // run until process end

  const numGoroutines = 1e5 // = 100,000. law of large numbers
  wg.Add(numGoroutines)
  before := memConsumed()  
  for i := numGoroutines; i > 0; i-- {
    go noop()
  }
  wg.Wait()
  after := memConsumed()

  fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
