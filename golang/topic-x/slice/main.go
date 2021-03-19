package main

import (
	"fmt"
	"runtime/debug"
)

func slpa(sl []int) {
	println(string(debug.Stack()))
	// fmt.Printf("addr sl %p\n", &sl)
	// fmt.Printf("addr sl arr %p\n", &sl[2])
}

func main() {
	a := []int{1,2,3}
	fmt.Printf("addr a %p\n", &a)
	fmt.Printf("addr a arr %p\n", &a[0])
	fmt.Printf("addr a arr %p\n", &a[1])
	fmt.Printf("addr a arr %p\n", &a[2])
	slpa(a)
}
