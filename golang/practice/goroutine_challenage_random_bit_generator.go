package main

import "fmt"

// Create a random bit generator that is a program that produces a sequence of 100000 randomly generated 1’s and 0’s using a goroutine.

func main() {

	ch := make(chan int)

	// consumer:
	go func() {
		for {
			fmt.Print(<-ch, " ")
		}
	}()
	for i := 0; i <= 100000; i++ {
		select {
			case ch <- 0:
			case ch <- 1:
		}
	}

}