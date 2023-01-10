package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"

	_ "embed"
)

//go:embed mock_client.tpl
var clientTpl string

type services []service

type service struct {
	Name string
}

func main() {
	svcs, err := svcs()
	if err != nil {
		panic(err)
	}

	fmt.Println(svcs)

	err = writeTemplate("test/client/mock_client.go", clientTpl, svcs)
	if err != nil {
		panic(err)
	}
}

func svcs() (services, error) {
	var (
		result = services{}
		set    = token.NewFileSet()
	)

	f, err := parser.ParseFile(set, "metal.go", nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		if gen, isGen := decl.(*ast.GenDecl); isGen {
			for _, spec := range gen.Specs {
				if t, isType := spec.(*ast.TypeSpec); isType {
					if t.Name.String() != "Client" {
						continue
					}

					if i, IsInterface := t.Type.(*ast.InterfaceType); IsInterface {
						for _, m := range i.Methods.List {
							for _, name := range m.Names {
								result = append(result, service{
									Name: name.Name,
								})
							}
						}
					}
				}
			}
		}
	}

	return result, nil
}

func writeTemplate(dest, text string, data any) error {
	t, err := template.New("").Funcs(sprig.FuncMap()).Parse(text)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return err
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	fmt.Printf("writing to %s\n", dest)

	return os.WriteFile(dest, p, 0755) // nolint:gosec
}
