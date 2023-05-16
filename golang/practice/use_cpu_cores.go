package main

import "runtime"

const NCPU = 4 // e.g. 4 for a quadcore processor

func DoAll() {
	sem := make(chan int, NCPU) // Buffering optional but sensible.
	for i := 0; i < NCPU; i++ {
		go DoPart(sem)
	}

	// Drain the channel sem, waiting for NCPU tasks to complete
	for i := 0; i < NCPU; i++ {
		<-sem // wait for one task to complete
	}

	// All done.
}

func DoPart(sem chan int) {
	// do the part of the computation
	sem <- 1 // signal that this piece is done
}

func main() {
	runtime.GOMAXPROCS(NCPU)
	DoAll()
}