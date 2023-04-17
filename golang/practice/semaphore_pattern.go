package main

import (
	"fmt"
	"time"
)

// An effecttive synchronization mechanism that let goroutines finish their execution first

type Semaphore interface {
	Acquire() // lock resource
	Release() // release resource
}

type semaphore struct {
	semC chan struct{}
}

func New(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}
func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}
func (s *semaphore) Release() {
	<-s.semC
}

/*
leverage the blocking behavior on Go buffered channel when channel is full by setting maxConcurrency parameter as the channel size.
When we calling Acquire(), the channel will be filled with an empty struct, this channel will be
blocking if it reaches its maximum value. And when we calling Release(), we take out the empty struct from the channel,
and the channel will be available for the next value and the channel will be unblocking.
*/

func main() {
	sem := New(3)
	doneC := make(chan bool, 1)
	totProcess := 10
	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totProcess {
				doneC <- true
			}
		}(i)
	}
	<-doneC
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second)
}
