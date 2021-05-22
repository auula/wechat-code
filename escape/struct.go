package main

import "fmt"

func main() {
	f := foo("Ding")
	fmt.Println(f)
}

type bar struct {
	s string
}

func foo(s string) *bar {
	f := new(bar)
	defer func() {
		f = nil
	}()
	f.s = s
	return f
}

// func foo(s string) bar {
// 	f := new(bar)
// 	defer func() {
// 		f = nil
// 	}()
// 	f.s = s
// 	return *f
// }
