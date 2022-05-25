package main

import (
	"fmt"
	"time"
)

func main() {
	/* after 1 second
	Cancelling doWork goroutine...
	doWork exited.
	Done.
	*/
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select { // Waiting...
				case s := <-strings: // nil channel
					fmt.Println(s)
				case <-done: // cancellation
					return
				}
			}

		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

func example2() {
	/*
		doWork start.
		red yellow blue ...........................................--Cancelling doWork goroutine--
		..
		doWork exited.
		Done.
	*/

	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("\ndoWork exited.")
			defer close(terminated)
			fmt.Println("doWork start.")
			for {
				select {
				case s := <-strings:
					if s == "" {
						fmt.Print(".")
					} else {
						fmt.Printf("%v ", s)
					}
				case <-done: // cancellation
					return
				}
			}

		}()
		return terminated
	}

	words := make(chan string, 3)
	words <- "red"
	words <- "yellow"
	words <- "blue"
	close(words)

	done := make(chan interface{})
	terminated := doWork(done, words)

	go func() {
		time.Sleep(50 * time.Microsecond)
		fmt.Println("--Cancelling doWork goroutine--")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

func example() {
	/*
		red
		yellow
		blue
		doWork exited.
		Done.
	*/
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	words := make(chan string, 3)
	words <- "red"
	words <- "yellow"
	words <- "blue"
	close(words)

	<-doWork(words)

	fmt.Println("Done.")
}

func leak() {
	/*
		Done.
	*/
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil) // goroutine leaks
	// <-doWork(nil) // fatal error: all goroutines are asleep - deadlock!

	fmt.Println("Done.")
}
