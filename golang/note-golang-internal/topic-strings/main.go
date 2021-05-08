package main

import (
	"fmt"
	"strings"
)

func main() {
	image := "dockerhub.com/project/service"

	fmt.Println(strings.SplitN(image, "/", 1))
	fmt.Println(strings.SplitN(image, "/", 2))
	fmt.Println(strings.Split(image, "/"))
	fmt.Println(strings.SplitAfterN(image, "/", 1))
	fmt.Println(strings.SplitAfterN(image, "/", 2))
	fmt.Println(strings.SplitAfter(image, "/"))
}
