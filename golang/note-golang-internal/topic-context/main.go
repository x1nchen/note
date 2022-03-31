package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancelRoot := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()
		println("root ctx done")
	}()

	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-childCtx.Done()
		println("child ctx done")
	}()

	// cancel root context will cancel child context
	cancelRoot()
	time.Sleep(3 * time.Second)
}
