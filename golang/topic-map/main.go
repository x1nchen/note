package main

import (
	"fmt"
)

func main() {
	m := make(map[int32]int32, 1)
	m[1] = 1
	v := m[1]
	fmt.Println(v)
}
