package main

import (
	"fmt"
	"time"
)

func main() {
	const row = 1024*64
	const col = 1024*64
	arr := [row][col]int{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			arr[i][j] = 0
		}
	}

	now := time.Now()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			_ = arr[i][j]
		}
	}

	fmt.Println("consume", time.Now().Sub(now))

	now = time.Now()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			_ = arr[j][i]
		}
	}
	fmt.Println("consume", time.Now().Sub(now))
}
