package main

import (
	"errors"
	"fmt"
)

func DemoReturnError() (err error) {
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
	}()
	return errors.New("new error")
}

func main() {
	err := DemoReturnError()
	fmt.Println(err)
}
