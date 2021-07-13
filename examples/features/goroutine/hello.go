package main

import "fmt"

func main() {
	fmt.Print("Start")

	go sayHello() // function

	go func () { // anonymous function
		fmt.Print("Hi")
	}()

	sayWelcome := func() { // variable + anonymous
		fmt.Print("Welcome")
	}
	go sayWelcome() // call variable

	for i := 1; i <= 100; i++ {
		fmt.Print(".")
	}

	fmt.Println("End")
}

func sayHello() {
	fmt.Print("Hello")
}

