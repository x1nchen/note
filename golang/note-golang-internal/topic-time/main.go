package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// year, mon, day := now.UTC().Date()
	// hour, min, sec := now.UTC().Clock()
	zone, offset := now.Zone()
	fmt.Println(zone, offset)

	lo, err := time.LoadLocation("Etc/GMT-1")
	if err != nil {
		panic(err)
	}

	fmt.Println(lo)

	nicosia, err := time.LoadLocation("Asia/Nicosia")
	if err != nil {
		panic(err)
	}

	t0 := time.Date(2021, 3,8, 0,0,0, 0, nicosia)
	fmt.Println("2021-3-08 00:00:00 ->",t0)

	t1 := time.Date(2021, 3,15, 0,0,0, 0, nicosia)
	fmt.Println("2021-3-15 00:00:00 ->", t1)

	t3 := time.Date(2021, 3,28, 2,55,0, 0, nicosia)
	fmt.Println("2021-3-28 2:55:00 ->", t3)

	t4 := time.Date(2021, 3,28, 3,0,0, 0, nicosia)
	fmt.Println("2021-3-28 3:00:00 ->",t4)

	t5 := time.Date(2021, 3,28, 4,0,0, 0, nicosia)
	fmt.Println("2021-3-28 4:00:00 ->", t5)
}

