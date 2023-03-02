## Channel

### Why create channel
Previously, the goroutines executed independently; they did not communicate. Of course, to be more useful, they have to communicate by sending and receiving information between them and, thereby, coordinating (synchronizing) their efforts. Goroutines could communicate by using , but this is  because this way of working introduces all the .

Instead, Go has a special type, the channel, which is , avoiding all the pitfalls of shared memory.
- channel guarantees synchronization
- The value of an uninitialized channel is nil.
- A channel can only transmit data-items of one datatype, e.g. chan int or chan string, but a channel can be made for any type, also custom types and for the empty interface{}. It is even possible (and sometimes useful) to create a channel of channels.
- A channel is, in fact, a typed message queue: data can be transmitted through it. It is a First In First Out (FIFO) structure