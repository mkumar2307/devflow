package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type FileSummary struct {
	Path      string
	Functions []string
	Structs   []string
}

func ParseFile(path string) (*FileSummary, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	summary := &FileSummary{
		Path: path,
	}

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			summary.Functions = append(summary.Functions, x.Name.Name)
		case *ast.TypeSpec:
			summary.Structs = append(summary.Structs, x.Name.Name)
		}
		return true
	})

	return summary, nil
}
