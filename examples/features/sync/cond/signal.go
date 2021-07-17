package main

import (
	"fmt"
	"sync"
)

func main() {
	simple()
	loop()
}

func loop() {
	c := sync.NewCond(&sync.Mutex{})
	var count = 5

	for i := 1; i <= count; i++ {
		go func(i int) {
			c.L.Lock()
			fmt.Printf("Hi %d\n", i)
			count--
			c.L.Unlock()
			c.Signal()
		}(i)
	}

	c.L.Lock()
	for count > 0 {
		c.Wait()
		fmt.Printf("Continue %d\n", count)
	}
	c.L.Unlock()
}

func simple() {
	c := sync.NewCond(&sync.Mutex{})
	var isTrue = false

	go func() {
		isTrue = true
		// fmt.Println("Simple")
		c.Signal()
	}()

	c.L.Lock()
	for isTrue == false {
		c.Wait()
	}
	c.L.Unlock()
}

/*
  Hi 5
  Hi 2
  Continue 3
  Hi 1
  Continue 2
  Hi 4
  Continue 1
  Hi 3
  Continue 0
*/
/*
  Hi 5
  Hi 3
  Hi 2
  Hi 1
  Hi 4
  Continue 0
*/
/*
  Hi 5
  Continue 4
  Hi 1
  Hi 3
  Hi 4
  Hi 2
  Continue 0
*/
