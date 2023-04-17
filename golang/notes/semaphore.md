## Overview
We can also use a channel for synchronization purposes, effectively using it as what is called a semaphore in traditional computing. To put it differently, to discover when a process (in a goroutine) is done, pass it a channel with which it can signal it is ready. A common idiom used to let the main program block indefinitely while other goroutines run is to place select {} as the last statement in the main function. This can also be done by using a channel to let the main program wait until the goroutine(s) complete(s), the so-called semaphore pattern.

From wikipedia:
> “Semaphore is a variable or abstract data type used to control access to a common resource by multiple processes in a concurrent system such as a multitasking operating system.”.

### Example

```go
done := make(chan bool)
// doSort is a lambda function, so a closure which knows the channel done:
doSort := func(s []int) {
  sort(s)
  done <- true
}
i := pivot(s)
go doSort(s[:i])
go doSort(s[i:])
<-done
<-done
```

### Semaphore for N goroutines

```go
type Empty interface {}
var empty Empty
...
data := make([]float64, N)
res := make([]float64, N)
sem := make(chan Empty, N) // semaphore
...
for i, xi := range data {
  go func (i int, xi float64) {
    res[i] = doSomething(i,xi)
    sem <- empty
  } (i, xi)
}
// wait for goroutines to finish
for i := 0; i < N; i++ { <-sem }
```

### mutex
Limit access to multiple resources, solve the reader-writers problem. This can be implemented by a buffered channel.
- the capacity of buffered channel is he number of resources we wish to synchronize
- the length (number of element currently stored) of the channel is the number of resources currently being used.
- the capacity minus the length of the channel is the number of resources currently available.

```go
type Empty interface {}
type semaphore chan Empty
sem = make(semaphore, N)
// acquire n resources
func (s semaphore) P(n int) {
  e := new(Empty)
  for i := 0; i < n; i++ {
    s <- e
  }
}
// release n resources
func (s semaphore) V(n int) {
  for i := 0; i < n; i++ {
    <-s
  }
}

/* mutexes */
func (s semaphore) Lock() {
  s.P(1)
}
func (s semaphore) Unlock() {
  s.V(1)
}
/* signal-wait */
func (s semaphore) Wait(n int) {
  s.P(n)
}
func (s semaphore) Signal() {
  s.V(n)
}
```