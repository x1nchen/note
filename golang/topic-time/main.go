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

	nicosia, err := time.LoadLocation("Asia/Nicosia")
	if err != nil {
		panic(err)
	}

	t := time.Date(2020, 3,15, 0,0,0, 0, nicosia)
	fmt.Println(t)

	// fmt.Println(now.In(time.UTC).Add(time.Hour*time.Duration(0)).Format("2006/01/02 15:04:05"))
	//
	// //date := time.Date(year, mon, day, hour, min, sec, 0, time.FixedZone("UTC", 1000))
	// //fmt.Println(date)
}

