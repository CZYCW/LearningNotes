## Background
- An applicaiton is a process running on a machine
- a processs is an independently executing entity that runs on its own space
- a processs is composed of one or more threads that are simultaneously executing entities that share the same address space. 
- Problem for multi-thread: race condition
  - solution: synchronize threads and lock the data

## Solution in Go
Communicating Sequential Process(CSP), also known as message passing-model -> Goroutine

- concurrently executing computations
- no one-to-one relationship between goroutines and threads
- a goroutine is mapped onto(multiplexed, executed by) one or more threads, according to their availability. 
- Goroutines run in the same address space
- goroutines should use channels to synchronize and communicate
- Goroutines run in the same address space
- The amount of available memory limits the number of goroutines
- goroutines run across multiple operational system threads
- can also run within threads, letting you handle a myriad of tasks with a relatively small memory footprint
- The stack of a goroutine grows and shrinks as needed.
- An executing goroutine can stop itself by calling `runtime.Goexit()`
- When one goroutine is very processor-intensive, you can call runtime.Gosched() periodically in your computation loops. -> This yields the processor,