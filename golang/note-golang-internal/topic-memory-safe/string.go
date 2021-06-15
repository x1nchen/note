package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	go func() {
		var i int = 0
		for {
			i = 1 - i
			m[i] = i
		}
	}()
	var j int = 0
	for {

		for {
			j = 1 - j
			x := m[j]
			fmt.Println(x)
		}
	}
}
