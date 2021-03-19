package main

import (
	"fmt"
	"runtime/debug"
)

type A struct{}

func slpa(a *A) {
	println(string(debug.Stack()))
	fmt.Printf("addr a  %p\n", a)
	fmt.Printf("addr of addr a %p\n", &a)
}

func main() {
	a := &A{}
	fmt.Printf("addr a %p\n", a)
	fmt.Printf("addr of addr a %p\n", &a)
	slpa(a)
}
