package main

import "fmt"

// func modify(s *[]int) {
// 	*s = append(*s, 4)
// 	fmt.Println(s)
// }

func temp(delete_funcs *[]func()) {
	var v = 1
	var b = 2
	var delete1 = func() {
		fmt.Println(v)
	}
	*delete_funcs = append(*delete_funcs, delete1)
	var delete2 = func() {
		fmt.Println(b)
	}
	*delete_funcs = append(*delete_funcs, delete2)
}

func main() {

	delete_funcs := []func(){}	
	temp(&delete_funcs)
	for _, f := range delete_funcs {
		f()
	}	

	// var s = []int{1, 2, 3}
	// modify(&s)
	// fmt.Println(s)

}