package main

import (
	"fmt"
)

func main() {
	// expr, _ := parser.ParseExpr("a * -1")
	// fmt.Printf("%#v\n", expr)
	fmt.Println("create before")
	i :=0
CREATE:
	fmt.Println("create after")
	if i != 10 {
		i++
		goto CREATE
	}
}
