package main

import (
  "fmt"
  "sync"
)

func main() {
  fmt.Println("Start")

  go sayHello() // function

  go func () { // anonymous function
    fmt.Println("Hi")
  }()

  var wg sync.WaitGroup
  sayWelcome := func() { // variable + anonymous
    defer wg.Done()
    fmt.Println("Welcome")
  }
  wg.Add(1)
  go sayWelcome() // call variable
  wg.Wait() // the join point

  for i := 1; i <= 30; i++ {
    fmt.Print(".")
  }

  fmt.Println("End")
}

func sayHello() {
  fmt.Println("Hello")
}
