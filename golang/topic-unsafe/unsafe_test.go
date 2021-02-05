package main

import (
	"testing"
	"unsafe"
)


// 获取 slice 的长度
// // runtime/slice.go
// type slice struct {
//    array unsafe.Pointer // 元素指针
//    len   int // 长度
//    cap   int // 容量
// }
//
// func makeslice(et *_type, len, cap int) slice
func TestGetSliceLength(t *testing.T) {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	// Len: &s => pointer => uintptr => pointer => *int => int
	t.Log(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	// Cap: &s => pointer => uintptr => pointer => *int => int
	t.Log(Cap, cap(s)) // 20 20
}


// 获取 map 的长度
// type hmap struct {
//	count     int
//	flags     uint8
//	B         uint8
//	noverflow uint16
//	hash0     uint32
//
//	buckets    unsafe.Pointer
//	oldbuckets unsafe.Pointer
//	nevacuate  uintptr
//
//	extra *mapextra
// }
// func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap
func TestGetMapLength(t *testing.T) {
	mp := make(map[int]int)
	mp[100] = 100
	mp[200] = 200

	count := **(**int)(unsafe.Pointer(&mp))
	t.Log(count, len(mp)) // 2 2


	flags := **(**uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&mp)) + uintptr(16)))
	t.Log(flags)
}