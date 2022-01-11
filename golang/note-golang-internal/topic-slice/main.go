package main

import (
	"fmt"
)

func main() {
	m := make(map[int][]int)

	m[1] = append(m[1], 1)
	fmt.Printf("%+v", m[1])

	IterateSlice(nil)
	fmt.Println("iterate")
}


func IterateSlice(targets []string) {
	for _, t := range targets {
		fmt.Println(t)
	}
}