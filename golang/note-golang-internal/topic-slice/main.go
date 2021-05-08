package main

import (
	"fmt"
)

func main() {
	m := make(map[int][]int)

	m[1] = append(m[1], 1)
	fmt.Printf("%+v", m[1])
}
