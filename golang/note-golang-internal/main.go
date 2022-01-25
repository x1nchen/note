package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/gops/agent"
)

var mu sync.RWMutex
var wg sync.WaitGroup
var a int

func Read() {
	mu.RLock()
	fmt.Println(a)
	mu.RUnlock()
}

func Write() {
	mu.Lock()
	defer mu.Unlock()
	PanicFunc()
}

func PanicFunc() {
	a = rand.Int()
	panic("panic func")
}

func main() {
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("read funcion call")
					Read()
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(ctx)
	}

	go func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				Write()
			}
		}
	}(ctx)

	wg.Wait()
	println("done")
}
