package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseDir(fset, "./", func(fi fs.FileInfo) bool {
		return fi.Name() != "main.go"
	}, 0)
	if err != nil {
		print("error: ", err.Error())
	} else {
		// visit all methods
		for _, pk := range f {
			println("package: ", pk.Name)

			ast.Inspect(pk, func(n ast.Node) bool {
				if fd, ok := n.(*ast.FuncDecl); ok {
					fmt.Println("Found function declaration: ", fd.Name)
				}
				return true
			})
		}
	}
}
