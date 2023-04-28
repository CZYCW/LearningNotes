## Ticker
time.Ticker, which is an object that repeatedly sends a time value on a contained channel C at a specified time interval.
```golang
type Ticker struct {
  C <-chan Time // the channel on which the ticks are delivered.
  // contains filtered or unexported fields
  ...
}
```

An example of limiting the process rate of request
```golang
import "time"

rate_per_sec := 10
var dur Duration = 1e9 / rate_per_sec
chRate := time.Tick(dur) // a tick every 1/10th of a second
for req := range requests {
  <- chRate // rate limit our Service.Method RPC calls
  go client.Call("Service.Method", req, ...)
}
```

### Timeout Pattern 1
```golang
timeout := make(chan bool, 1)
go func() {
  time.Sleep(1e9) // one second
  timeout <- true
}()

select {
  case <-ch:
  // a read from ch has occurredk
  case <-timeout:
  break
}
```

### Timeout Pattern 2
```golang
ch := make(chan error, 1)
go func() { ch <- client.Call("Service.Method", args, &reply) } ()
  select {
    case resp := <-ch:
      // use resp and reply
    case <-time.After(timeoutNs):
     // call timed out
   break
}
```

### Timeout Pattern 3
Suppose we have a program that reads from multiple replicated databases simultaneously. The program needs only one of the answers, and it should accept the answer that arrives first. The function Query takes a slice of database connections and a query string. It queries each of the databases in parallel and returns the first response it receives:

```golang
func Query(conns []Conn, query string) Result {
  ch := make(chan Result, 1)
  for _, conn := range conns {
    go func(c Conn) {
      select {
        case ch <- c.DoQuery(query):
          default:
      }
    }(conn)
  }
return <- ch
}
```