package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	filename := "print_name.go"

	f, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		print("error: ", err.Error())
	} else {
		ast.Print(fset, f)
	}
}
