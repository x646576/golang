package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int // nil

	select {
	case <-c: // read nil channel
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}

/*
Timed out.
*/
