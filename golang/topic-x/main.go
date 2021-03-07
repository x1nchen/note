package main

import (
	"runtime/debug"
)

func slpa(sl []int) {
	println(string(debug.Stack()))
}

func main() {
	slpa([]int{1,2,3})
}



