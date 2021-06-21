package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		c1 := make(chan os.Signal, 1)
		signal.Notify(c1, os.Interrupt)
		// Block until a signal is received.
		s1 := <-c1
		fmt.Println("Got signal s1:", s1)

	}()

	// Block until a signal is received.
	s := <-c

	fmt.Println("Got signal s:", s)
	time.Sleep(time.Second*10)
}

