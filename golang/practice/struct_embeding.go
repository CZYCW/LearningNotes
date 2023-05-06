package main
import "fmt"

type Base struct {
	b int
}

type Container struct { // Container is the embedding struct
	Base // Base is the embedded struct
	c    string
}

func main() {
	co := Container{}
	co.b = 1
	co.c = "string"
	fmt.Printf("co -> {b: %v, c: %v}\n", co.b, co.c)

}
