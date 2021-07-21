package main

import (
	"fmt"
	"sync"
)

func main() {
	// read and write
	var rwStream chan interface{}     // declare. empty interface type
	rwStream = make(chan interface{}) // instantiate

	readStream := make(<-chan interface{})  // read only
	writeStream := make(chan<- interface{}) // write only

	/*
		<-writeStream
		invalid operation: <-writeStream (receive from send-only type chan<- interface {})

		readStream <- struct{}{}
		invalid operation: readStream <- struct {}{} (send to receive-only type <-chan interface {})
	*/

	// implicitly convert bidirectional channels to unidirectional channels
	var receiveChan <-chan interface{}
	var sendChan chan<- interface{}
	dataStream := make(chan interface{})

	receiveChan = dataStream
	sendChan = dataStream

	// stdout <- stream <- "Hello"
	stringStream := make(chan string)
	go func() {
		/* fatal error: all goroutines are asleep - deadlock!
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

	// stricter type
	intStream := make(chan int)
	close(intStream) // closed channel
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer) // (false): 0

	numberStream := make(chan int)
	go func() {
		defer close(numberStream) // ensure that the channel is closed
		for i := 1; i <= 5; i++ {
			numberStream <- i
		}
	}()

	// automatically break the loop when a channel is closed
	// range does not return the second boolean value
	for integer := range numberStream {
		fmt.Printf("%v ", integer) // 1 2 3 4 5
	}
	fmt.Println()

	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin // waiting...
			fmt.Printf("%v has begun\n", i)
		}(i)
	}
	fmt.Println("Unblocking goroutines...")
	close(begin) // unblock all goroutines
	wg.Wait()

	var bufferedStream chan interface{}
	bufferedStream = make(chan interface{}, 4)    // 4 capacity
	unbufferedStream := make(chan interface{}, 0) // 0 capacity = make(chan interface{})

	// pass
	pass(rwStream, readStream, writeStream, receiveChan, sendChan)
}

func pass(v ...interface{}) {}
