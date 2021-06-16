package main

import (
	"fmt"
	"time"
)

const (
	FIRST  = "hello world"
	SECOND = "abcde"
)

func main() {
	m := ""
	go func() {
		var i int = 1
		for {
			i = 1 - i
			if i == 0 {
				m = FIRST
			} else {
				m = SECOND
			}
			time.Sleep(10)
		}
	}()
	for {
		if m == "hello" {
			panic(m)
		}
		fmt.Println(m)
		time.Sleep(10)
	}
}
