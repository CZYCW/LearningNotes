package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go receiver(ch)
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent")
}

func receiver(ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
	fmt.Println("received")
}
