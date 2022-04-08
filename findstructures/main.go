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
				if gd, ok := n.(*ast.GenDecl); ok {
					for _, spec := range gd.Specs {
						switch spec.(type) {
						case *ast.TypeSpec:
							typeSpec := spec.(*ast.TypeSpec)

							switch typeSpec.Type.(type) {
							case *ast.StructType:
								println("struct: ", typeSpec.Name.Name)

								structType := typeSpec.Type.(*ast.StructType)

								for _, field := range structType.Fields.List {
									i := field.Type.(*ast.Ident)
									fieldType := i.Name

									for _, name := range field.Names {
										fmt.Printf("\tField: name=%s type=%s\n", name.Name, fieldType)
									}

								}
							}
						}
					}
				}
				return true
			})
		}
	}
}
