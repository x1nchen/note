package main

import (
	"sync/atomic"
)

// func fn(m map[int]int) {
// 	println(string(debug.Stack()))
// 	m = make(map[int]int)
// }
//
// func main() {
// 	var m map[int]int
// 	fmt.Printf("%p", m)
// 	fn(m)
// 	fmt.Println(m == nil)
// }


func main() {
	atomic.CompareAndSwapInt64()
}
