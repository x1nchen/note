package main

import (
	"fmt"
	"reflect"
)

type  A  struct {
	a int
}

type B struct {
	b int
}


func main() {
	ai1 := &A{1}
	ai2 := &A{1}

	fmt.Println(reflect.DeepEqual(ai1, ai2))
}