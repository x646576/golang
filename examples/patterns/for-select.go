package main

import "fmt"

func main() {
	iterate()
	waiting()
}

func iterate() {
	fmt.Println("Sending iteration variables out on a channel")

	var done chan interface{}
	stringStream := make(chan interface{}, 3)

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}

	close(stringStream)

	for c := range stringStream {
		fmt.Printf("%v ", c)
	}
	fmt.Println("\n")
}

func waiting() {
	fmt.Println("Looping infinitely waiting to be stopped")

	done := make(chan interface{})

	go func() {
		close(done)
	}()

	for {
		select {
		case <-done:
			fmt.Println("\n")
			return
		default:
			fmt.Print(".") // Do non-preemptable work
		}
		// or here
		fmt.Print(" ") // Do non-preemptable work
	}
}
