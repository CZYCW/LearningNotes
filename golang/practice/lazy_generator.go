package main
import (
"fmt"
)

var resume chan int

// a function that create a channel and a goroutine to constantly send integers to the channel
func integers() chan int {
  yield := make (chan int)
  count := 0
  go func () {
    for {
      yield <- count
      count++
    }
  } ()
  return yield
}

func generateInteger() int {
  return <-resume
}

func main() {
  resume = integers()
  fmt.Println(generateInteger()) //=> 0
  fmt.Println(generateInteger()) //=> 1
  fmt.Println(generateInteger()) //=> 2
}