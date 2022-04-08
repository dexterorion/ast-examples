package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"github.com/dexterorion/ast-examples/buildast/specs"
)

var (
	BaseGenFolder = "generation/"
)

func main() {
	// fset := token.NewFileSet()
	// f := fset.AddFile("build.go", 1, 100)

	// err := ast.Fprint(os.Stdout, fset, f, nil)
	// if err != nil {
	// 	panic(err)
	// }

	spec, err := specs.NewSpecFromFile("./specification.json")
	if err != nil {
		panic(err)
	}

	generate(spec)
	guarantee(spec)
}

func generate(spec *specs.Spec) {
	err := os.Mkdir(BaseGenFolder+spec.Package, os.ModePerm)
	if err != nil {
		panic(err)
	}

	for _, f := range spec.Files {
		file, err := os.Create(BaseGenFolder + spec.Package + "/" + f.Name)
		if err != nil {
			println(err.Error())
			continue
		}

		w := bufio.NewWriter(file)

		w.WriteString("package " + spec.Package + "\n\n")

		for _, str := range f.Structures {
			w.WriteString("type " + str.Name + " struct {\n")

			for _, field := range str.Fields {
				w.WriteString("\t" + strings.Title(field.Name) + " " + field.Type + " `json:\"" + strings.ToLower(field.Name) + "\"`\n")
			}
			w.WriteString("}\n\n")
		}

		w.Flush()
	}
}

func guarantee(spec *specs.Spec) {
	fset := token.NewFileSet()

	f, err := parser.ParseDir(fset, BaseGenFolder+"/"+spec.Package, nil, 0)
	if err != nil {
		print("error: ", err.Error())
	}

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

func getters(spec *specs.Spec) {

}
