package main

import (
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	filename := "print.go"

	_, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		print("error : ", err.Error())
	} else {
		print("everything is fine")
	}
}
