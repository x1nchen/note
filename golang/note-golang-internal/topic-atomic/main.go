package main

import (
	"sync/atomic"
)

func main() {
	var a uint64 = 1
	println(atomic.AddUint64(&a, 1))
	println(a)
}