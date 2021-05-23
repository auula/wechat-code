package main

import "fmt"

func main() {
	f := foo("Ding")
	// ptr := &f
	// fmt.Printf("main() %p \n", *ptr)
	fmt.Println(f)
}

type bar struct {
	s string
}

// func foo(s string) *bar {
// 	f := new(bar)
// 	ptr := &f
// 	defer func() {
// 		//fmt.Printf("foo().func() %p \n", ptr)
// 		*ptr = nil
// 		//fmt.Printf("foo().func() %p \n", ptr)
// 	}()
// 	f.s = s
// 	//fmt.Printf("foo() %p \n", ptr)
// 	return *ptr
// }

func foo(s string) bar {
	f := new(bar)
	ptr := &f
	defer func() {
		//fmt.Printf("foo().func() %p \n", ptr)
		f = nil
	}()
	f.s = s
	//fmt.Printf("foo() %p \n", ptr)
	return **ptr
}
