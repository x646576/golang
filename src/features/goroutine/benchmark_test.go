// go test -bench=. -cpu=1 benchmark_test.go

package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // wait until begin
		for i := 0; i < b.N; i++ {
			c <- token // send messages
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // wait until begin
		for i := 0; i < b.N; i++ {
			<-c // receive messages
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin) // tell the 2 goroutines to begin
	wg.Wait()
}
