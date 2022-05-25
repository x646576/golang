package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// unbufferedStream := make(chan interface{})
	// unbufferedStream := make(chan interface{}, 0)

	// bufferedStream := make(chan interface{}, 4)
	var bufferedStream chan interface{}
	bufferedStream = make(chan interface{}, 4) // 4 capacity

	var stdoutBuff bytes.Buffer // in-memory buffer. little faster than stdout. not guarantee.
	defer stdoutBuff.WriteTo(os.Stdout)

	go func() {
		defer close(bufferedStream)
		defer fmt.Fprintln(&stdoutBuff, "  Producer Done.")
		for i := 0; i < 4; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			bufferedStream <- i
		}
	}()

	for integer := range bufferedStream {
		fmt.Fprintf(&stdoutBuff, "Received: %v\n", integer)
	}
}

/*
Sending: 0
Sending: 1
Sending: 2
Sending: 3
  Producer Done.
Received: 0
Received: 1
Received: 2
Received: 3
*/
