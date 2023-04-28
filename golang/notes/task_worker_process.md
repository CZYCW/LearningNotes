### Overview
Suppose we have to perform several tasks; a task is performed by a worker (process). A Task can be defined as a struct (the concrete details are not important here):

```golang
type Task struct {
  // some state
}
```

### Paradigm 1: shared memory to synchronize
The pool of tasks is shared memory. To synchronize the work and to avoid race conditions, we have to guard the pool with a Mutex lock:
```golang
type Pool struct {
  Mu sync.Mutex
  Tasks []Task
}
```

A worker locks the pool, takes the first task from the pool, unlocks the pool, and then processes the task. The lock guarantees that only one worker process at a time can access the pool. 
```go
func Worker(pool *Pool) {
  for {
    pool.Mu.Lock()
    // begin critical section:
    task := pool.Tasks[0] // take the first task
    pool.Tasks = pool.Tasks[1:] // update the pool of tasks
    // end critical section
    pool.Mu.Unlock()
    process(task)
  }
}
```
 ### Paradigm 2: channels to synchronize

 ```go
 func main() {
  pending, done := make(chan *Task), make(chan *Task)
  go sendWork(pending) // put tasks to do on the channel
  for i := 0; i < N; i++ { // start N goroutines to do work
    go Worker(pending, done)
  }
  consumeWork(done) // continue with the processed tasks
}
```

worker
```go
func Worker(in, out chan *Task) {
  for {
    t := <-in
    process(t)
    out <- t
  }
}
```