package main

import (
	"fmt"
)

func main() {
	a := []int{1,2,3}
	a = append([]int{1,2}, a...)
	fmt.Println(a)
}
