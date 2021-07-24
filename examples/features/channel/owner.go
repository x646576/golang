package main

import "fmt"

func main() {
	chanOwner := func() <-chan int { // Ecapsulate. via a reader channel.
		resultStream := make(chan int, 5) // Instantiate.
		go func() {
			defer close(resultStream) // Close.
			for i := 0; i <= 5; i++ {
				resultStream <- i // Writes. or pass ownership to another goroutine.
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()

	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

/*
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Done receiving!
*/
