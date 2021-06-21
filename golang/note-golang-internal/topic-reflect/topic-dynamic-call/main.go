package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(handle(i))
	}
}

func handle(i int) (x int) {
	defer func(i int) {
		if err := recover(); err != nil {
			x = i
		} else {
			x = i
		}
	}(i)

	if i == 1 {
		i = i / (i - 1)
	}

	return
}
