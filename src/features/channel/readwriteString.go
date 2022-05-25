package main

import "fmt"

func main() {
	// stdout <- stream <- "Hello"
	stringStream := make(chan string)
	go func() {
		/* make nil channel
		fatal error: all goroutines are asleep - deadlock!
		if 0 != 1 {
			return
		}
		*/
		stringStream <- "Hello channels!" // will not exit until the write is successful
		stringStream <- "Hi, there."
	}()
	fmt.Println(<-stringStream) // wait until a value is placed on the channel

	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation) // (true): Hi, there.
}

/*
Hello channels!
(true): Hi, there.
*/
