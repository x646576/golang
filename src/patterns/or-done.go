package main

import "fmt"

func main() {

	total := 5
	breakPoint := 3

	myChan := make(chan interface{})
	go func() {
		defer close(myChan)
		for i := 1; i <= total; i++ {
			myChan <- i
			if i == breakPoint {
				return
			}
		}
	}()

	done := make(chan interface{})
	defer close(done)

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	fmt.Printf("Loop %v\n", total)
	for val := range orDone(done, myChan) {
		fmt.Printf("%v / %v\n", val, total)
	}
	fmt.Println("End")

}

/*
Loop 5
1 / 5
2 / 5
3 / 5
End
*/
