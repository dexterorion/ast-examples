package main

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

type Structure struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name string
	Type string
}

func main() {
	mapStructure := map[string]*Structure{}

	fset := token.NewFileSet()

	src, err := ioutil.ReadFile("model.go")
	if err != nil {
		panic(err)
	}

	f, err := parser.ParseFile(fset, "model.go", nil, 0)
	if err != nil {
		print("error: ", err.Error())
	} else {
		ast.Inspect(f, func(n ast.Node) bool {
			if gd, ok := n.(*ast.GenDecl); ok {
				for _, spec := range gd.Specs {
					switch spec.(type) {
					case *ast.TypeSpec:
						typeSpec := spec.(*ast.TypeSpec)

						switch typeSpec.Type.(type) {
						case *ast.StructType:
							structure := &Structure{
								Name: typeSpec.Name.Name,
							}

							structType := typeSpec.Type.(*ast.StructType)

							structure.Fields = make([]*Field, 0, len(structType.Fields.List))
							for _, field := range structType.Fields.List {
								fName := field.Names[0].Name
								fType := string(src[field.Type.Pos()-1 : field.Type.End()-1])

								structure.Fields = append(structure.Fields, &Field{
									Name: fName,
									Type: fType,
								})
							}

							mapStructure[typeSpec.Name.Name] = structure
						}
					}
				}
			}
			return true
		})
	}

	// generate getters and setters
	os.Remove("model_methods.go")

	fileMethods, err := os.Create("model_methods.go")
	if err != nil {
		panic(err)
	}

	buf := bufio.NewWriter(fileMethods)

	buf.WriteString("package " + f.Name.Name + "\n\n")

	for _, structure := range mapStructure {
		for _, field := range structure.Fields {
			buf.WriteString("func (s *" + structure.Name + ") Set" + strings.Title(field.Name) + "(" + strings.ToLower(field.Name) + " " + field.Type + ") {\n")
			buf.WriteString("\ts." + field.Name + " = " + strings.ToLower(field.Name) + "\n")
			buf.WriteString("}\n\n")
		}
	}

	buf.Flush()
}
