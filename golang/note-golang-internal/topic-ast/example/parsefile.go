// parsefile.go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	// if the src parameter is nil, then will auto read the second filepath file
	f, _ := parser.ParseFile(fset, "example.go", src, parser.Mode(0))

	for _, d := range f.Decls {
		ast.Print(fset, d)
		fmt.Println()
	}

	for _, d := range f.Imports {
		ast.Print(fset, d)
		fmt.Println()
	}

	ast.Print(fset, f.Name)

}

var src = `package pppppp
import _ "log"
func add(n, m int) {}
`