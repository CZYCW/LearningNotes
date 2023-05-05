package main
import (
"fmt"
)

type Any interface{} // defines an alias Any for the empty interface
type EvalFunc func(Any) (Any, Any) // defines a function type EvalFunc that takes an Any and returns a tuple of two Any’s

func main() {
	// an anonymous function with type EvalFunc is defined and assigned to the name evenFunc
    evenFunc := func(state Any) (Any, Any) {
        os := state.(int)
        ns := os + 2
        return os, ns
    }

    even := BuildLazyIntEvaluator(evenFunc, 0)
    for i := 0; i < 10; i++ {
        fmt.Printf("%vth even: %v\n", i, even())
    }
}

// the generalized version of BuildLazyIntEvaluator, which works for any type
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
    retValChan := make(chan Any)
    loopFunc := func() {
        var actState Any = initState
        var retVal Any
        for {
            retVal, actState = evalFunc(actState)
            retValChan <- retVal
        }
    }
    retFunc := func() Any {
        return <-retValChan
    }
    go loopFunc()
    return retFunc
}

// It takes a function of type EvalFunc as a parameter, as well as a variable of any type initState
// Its return value is a function that has no parameters and returns an int
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
    ef := BuildLazyEvaluator(evalFunc, initState)
    return func() int {
        return ef().(int)
    }
}