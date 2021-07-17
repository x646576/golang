package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("[Start]")

	go sayHello() // function

	go func() { // anonymous function
		fmt.Println("  2. Hi")
	}()

	var wg sync.WaitGroup
	sayWelcome := func() { // variable + anonymous
		defer wg.Done()
		fmt.Println("    3. Welcome")
	}
	wg.Add(1)
	go sayWelcome() // call variable
	wg.Wait()       // the join point

	// var wg sync.WaitGroup
	salutation := "        4. hello (before Wait)"
	wg.Add(1)
	go func() { // closure
		defer wg.Done()
		salutation = "        4. welcome (after Wait)"
	}()
	fmt.Println(salutation) // "hello"
	wg.Wait()
	fmt.Println(salutation) // "welcome"

	for _, salutation := range []string{"          5. hello", "          5. greetings", "          5. good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation) // "good day" "good day" "good day"
		}()
	}
	wg.Wait()

	for _, salutation := range []string{"            6. hello", "            6. greetings", "            6. good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation) // "good day" "greetings" "hello"
		}(salutation) // copy variable
	}
	wg.Wait()

	fmt.Println("[End]")
}

func sayHello() {
	fmt.Println("1. Hello")
}
