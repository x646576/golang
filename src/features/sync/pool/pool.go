package main

import (
	"fmt"
	"sync"
)

func main() {

	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // "Creating new instance."
	instance := myPool.Get() // "Creating new instance."
	myPool.Put(instance)
	myPool.Get()

}
