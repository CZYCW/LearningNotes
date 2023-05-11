## Context
Sometimes, you know you need to compute a value before you actually need to use the value. In this case, you can potentially start computing the value on another processor and have it ready when you need it. Futures are easy to implement via closures and goroutines, and the idea is similar to generators, except a future needs only to return one value.

### Example
When we are calculating the inverse of the product of 2 matrics, we can start calculating the inverse of each matrix on a different processor, and then multiply the results when they are both ready.

Original Version without parallelization:
```go
func InverseProduct(a Matrix, b Matrix) {
  a_inv := Inverse(a)
  b_inv := Inverse(b)
  return Product(a_inv, b_inv)
}
```

After parallelization:
- InverseFuture() launches a closure as a goroutine, which puts the resultant inverse matrix on a channel future, as a result:
```go
func InverseProduct(a Matrix, b Matrix) {
  a_inv_future := InverseFuture(a) // started as a goroutine
  b_inv_future := InverseFuture(b) // started as a goroutine
  a_inv := <-a_inv_future
  b_inv := <-b_inv_future
  return Product(a_inv, b_inv)
}

func InverseFuture(a Matrix) (chan Matrix) {
  future := make(chan Matrix)
  go func() { future <- Inverse(a) }()
  return future
}
```

### Example for API Implementation
```go
// futures used internally
type futureMatrix chan Matrix;

// API remains the same
func Inverse (a Matrix) Matrix {
    return <-InverseAsync(promise(a))
}

func Product (a Matrix, b Matrix) Matrix {
    return <-ProductAsync(promise(a), promise(b))
}

// expose async version of the API
func InverseAsync (a futureMatrix) futureMatrix {
    c := make (futureMatrix);
    go func () { c <- inverse(<-a) } ();
    return c
}

func ProductAsync (a futureMatrix) futureMatrix {
    c := make (futureMatrix);
    go func () { c <- product(<-a) } ();
    return c
}

// actual implementation is the same as before
func product (a Matrix, b Matrix) Matrix {
    ....
}

func inverse (a Matrix) Matrix {
    ....
}

// utility fxn: create a futureMatrix from a given matrix
func promise (a Matrix) futureMatrix {
    future := make (futureMatrix, 1);
    future <- a;
    return future;
}
```

Actual usage of the implementation:
```go
func InverseProduct (a Matrix, b Matrix) {
    a_inv := Inverse(a);
    b_inv := Inverse(b);
    return Product(a_inv, b_inv);
}
```