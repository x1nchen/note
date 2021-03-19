package main

import (
	"fmt"
	"reflect"
)

func main() {
	Fuck(nil)
}

func Fuck(a interface{}) {
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a).Kind())
}
