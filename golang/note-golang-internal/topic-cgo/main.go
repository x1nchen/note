package main

// #include <stdio.h>
// #include <stdlib.h>
// void myFunc(char *msg) {
// printf("before Hello %s,%p!\n", msg,msg);
// msg = "999";
// printf("after Hello %s,%p!\n", msg,msg);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	m := C.CString("world")
	fmt.Printf("msg.msg:%p,%p\n",unsafe.Pointer(m),m)
	C.free(unsafe.Pointer(m))
	C.myFunc(m)
}