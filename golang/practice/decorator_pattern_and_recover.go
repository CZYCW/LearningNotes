package main

import "fmt"

type HandleLogic func(int)

func Handle(handleLogic HandleLogic, i int) {
	func(int) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("recover")
			}
		}()
		handleLogic(i)
	}(i)
}
func handleLogic(i int) {
	fmt.Println("handleLogic")
	panic("panic")
}
func main() {
	go Handle(handleLogic, 1)
	fmt.Print("mains")
}
