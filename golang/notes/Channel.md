## Channel

### Why create channel
Previously, the goroutines executed independently; they did not communicate. Of course, to be more useful, they have to communicate by sending and receiving information between them and, thereby, coordinating (synchronizing) their efforts. Goroutines could communicate by using , but this is  because this way of working introduces all the .

Instead, Go has a special type, the channel, which is , avoiding all the pitfalls of shared memory.
- channel guarantees synchronization
- The value of an uninitialized channel is nil.
- only one goroutine has access to a data item at any given time: so data races cannot occur, by design
- A channel can only transmit data-items of one datatype, e.g. chan int or chan string, but a channel can be made for any type, also custom types and for the empty interface{}. It is even possible (and sometimes useful) to create a channel of channels.
- A channel is, in fact, a typed message queue: data can be transmitted through it. It is a First In First Out (FIFO) structure
- Channels are first-class objects. They can be stored in variables, passed as arguments to functions, returned from functions, and sent themselves over channels. Moreover, they are typed, allowing the type system to catch programming errors like trying to send a pointer over a channel of integers.''
  
### Blocking of goroutines
By default, communication is synchronous, and unbuffered, which means the send operation does not complete until there is a receiver to accept the value. 
- A send operation on a channel (and the goroutine or function that contains it) blocks until a receiver is available for the same channel ch. If there’s no recipient for the value on ch, no other value can be put in the channel, which means no new value can be sent in ch when the channel is not empty. So, the send operation will wait until ch becomes available again. This is the case when the channel-value is received (can be put in a variable).
- A receive operation for a channel blocks (and the goroutine or function that contains it) until a sender is available for the same channel. If there is no value in the channel, the receiver blocks. Although this seems a severe restriction, this offers a simple form of synchronizing and which works well in most practical situations.

### Buffered channel
- A buffered channel is a channel that can hold a fixed number of values. A buffered channel has two parameters: the type of data it holds and the number of elements it can hold.

```go
buf := 100
ch1 := make(chan string, buf)
```can 