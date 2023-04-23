package main

// Getting the values out of different, concurrently executing goroutines can be accomplished with the select keyword, 
// which closely resembles the switch control statement, and is sometimes called the communications switch.

import (
"fmt"
"time"
"runtime"
)


/**
 What select does is,

If all are blocked, it waits until one can proceed.
When none of the channel operations can proceed, and the default clause is present, then this is executed because the default is always runnable, or ready to execute.
.
If multiple can proceed, one is chosen at random
**/

func main() {
    runtime.GOMAXPROCS(2) 
    ch1 := make(chan int)
    ch2 := make(chan int)
    go pump1(ch1)
    go pump2(ch2)
    go suck(ch1, ch2)
    time.Sleep(1e9)
}

func pump1(ch chan int) {
    for i:=0; ; i++ {
        ch <- i*2
    }
}

func pump2(ch chan int) {
    for i:=0; ; i++ {
        ch <- i+5
    }
}

func suck(ch1 chan int,ch2 chan int) {
    for {
        select {

            case v:= <- ch1:
                fmt.Printf("Received on channel 1: %d\n", v)
            case v:= <- ch2:
                fmt.Printf("Received on channel 2: %d\n", v)
        }
    }
}