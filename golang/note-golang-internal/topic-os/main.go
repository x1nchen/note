package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	binfile := os.Args[0]
	log.Println(binfile)

	Executable, err := os.Executable()
	log.Println(Executable, err)

	tmps := strings.Split(Executable, "/")
	log.Println(tmps[len(tmps)-1])

	// 2022/01/17 14:39:06 ./topic-os
	// 2022/01/17 14:39:06 /Users/leon/repo/github.com/x1nchen/note/golang/note-golang-internal/topic-os/topic-os <nil>
	// 2022/01/17 14:39:06 topic-os

	log.Println(filepath.Base(os.Args[0]))
	log.Println(filepath.Base(Executable))
}
