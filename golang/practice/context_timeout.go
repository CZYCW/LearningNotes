package main

import (
	// "context"
	"fmt"
	"time"
)

func main() {
	// Channel used to receive the result from doSomething function
	// ch := make(chan string, 1)

	// Create a context with a timeout of 5 seconds
	// ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	// Start the doSomething function
	go doSomething()
	fmt.Print("Waiting for doSomething wakeup.. \n")
	// select {
	// case <-ctxTimeout.Done():
	// 	fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	// case result := <-ch:
	// 	fmt.Printf("Received: %s\n", result)
	// }
	return 
}

func doSomething() {
	fmt.Println("doSomething Sleeping...")
	time.Sleep(time.Second * 5)
	fmt.Println("doSomething Wake up...")
}