package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	select {
		case <-ctx.Done():
			fmt.Println("ctx done")
		default:
			fmt.Println("default")
	}

	<-ctx.Done()

	// drain out channel
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i ++ {
			go func() {
				for _  = range ch {
					ccc
				}
			}()
		}
	}()

	ch <- 1
	close(ch)

}
