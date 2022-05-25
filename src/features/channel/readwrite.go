package main

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

	// pass
	pass(rwStream, readStream, writeStream, receiveChan, sendChan)
}

func pass(v ...interface{}) {}
