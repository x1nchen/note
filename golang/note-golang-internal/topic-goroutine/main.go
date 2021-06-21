package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}

	var ch = make(chan int)
	<- ch
}