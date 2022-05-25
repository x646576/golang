package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for { // keep repeating
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	defer close(done)

	for v := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", v) // 1 1 1 1 / 1 1 1 1 / 1 1
	}
	fmt.Println()

	// repeat function

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	rand := func() interface{} {
		return rand.Int()
	}

	for v := range take(done, repeatFn(done, rand), 10) {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	/*
		5577006791947779410 8674665223082153551 6129484611666145821
		4037200794235010051 3916589616287113937 6334824724549167320
		605394647632969758 1443635317331776148 894385949183117216
		2775422040480279449
	*/

	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s\n", message) // message: Iam.Iam.I...

}
