package main

import (
	"fmt"
	"sync"
)

func main() {
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
}

/*
Unblocking goroutines...
1 has begun
4 has begun
0 has begun
3 has begun
2 has begun
*/
