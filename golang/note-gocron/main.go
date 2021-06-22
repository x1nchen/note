package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	ex := make(chan int)
	c := cron.New()
	_, err := c.AddFunc("41 11 24 3 *", func() {
		fmt.Println("trigger and exit")
		ex <- 1
	})
	if err != nil {
		panic(err)
	}
	c.Start()

	<-ex
}
