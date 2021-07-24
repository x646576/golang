package main

import "fmt"

func main() {
	var value interface{}
	var ok interface{}

	// int default 0
	intStream := make(chan int) // stricter type
	close(intStream)            // closed channel
	value, ok = <-intStream
	fmt.Printf("(%v): %v\n", ok, value) // (false): 0

	// string default ""
	stringStream := make(chan string) // stricter type
	close(stringStream)               // closed channel
	value, ok = <-stringStream
	fmt.Printf("(%v): \"%v\"\n", ok, value) // (false): ""

	numStream := make(chan int, 1) // stricter type, buffer 1
	numStream <- 99
	value, ok = <-numStream
	fmt.Printf("(%v): %v\n", ok, value) // (true): 99
	close(numStream)                    // closed channel
	value, ok = <-numStream
	fmt.Printf("(%v): %v\n", ok, value) // (false): 0
}

/*
(false): 0
(false): ""
(true): 99
(false): 0
*/
