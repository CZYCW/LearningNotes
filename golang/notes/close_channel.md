## Closing a Channel
Closing a channel is only necessary when the receiver must be told that there are no more values coming. Only the sender should close a channel, never the receiver.

### Signaling
```go
ch := make(chan int)
defer close(ch)
```
- marks channel as unable to receive more values. 
- Sending to or closing a closed channel causes a run-time panic

### Detection
```go
v, ok := <-ch // ok is true if v received a value
```

a better way
```go
if v, ok := <-ch; ok {
  ...
}
```

a forloop
```go
v, ok := <-ch
if !ok {
  break
}
// process(v)
```

### Example
```go
package main
import "fmt"

func main() {
  ch := make(chan string)
  go sendData(ch)
  getData(ch)
}

func sendData(ch chan string) {
  ch <- "Washington"
  ch <- "Tripoli"
  ch <- "London"
  ch <- "Beijing"
  ch <- "Tokyo"
  close(ch)
}

func getData(ch chan string) {
  for {
    input, open := <-ch
    if !open {
      break
    }
    fmt.Printf("%s ", input)
  }
}
```

### another way
It is even better to practice reading the channel with a for-range statement because this will automatically detect when the channel is closed:
```go
for input := range ch {
  Process(input)
}
```

