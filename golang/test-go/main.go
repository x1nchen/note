package main

import (
	"encoding/json"
)

func main() {
	if err := json.Unmarshal([]byte{}, nil); err == nil {
		panic("not err")
	} else {
		panic(err)
	}
}
