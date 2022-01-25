package main

import (
	"os"
)

func main() {
	process, err := os.Executable()
	if err != nil {
		panic(err)
	}
	println(process)
}
