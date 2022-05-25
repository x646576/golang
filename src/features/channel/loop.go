package main

import "fmt"

func main() {
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
}

/*
1 2 3 4 5
*/
