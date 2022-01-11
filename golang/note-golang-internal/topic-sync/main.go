package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		fmt.Println("exit")
	}()

	go func() {
		for i := 0; i<10; i++ {
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(5*time.Second)
}
