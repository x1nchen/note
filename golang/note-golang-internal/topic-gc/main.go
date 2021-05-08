package main

var A WB
var B WB

type WB struct {
	Obj *int
}

func simpleSet(c *int) {
	A.Obj = nil
	B.Obj = c
	// if GC Begin
	// Scan A
	A.Obj = c
	B.Obj = nil
	// scan B
}

func main() {
	i := 1
	simpleSet(&i)
}