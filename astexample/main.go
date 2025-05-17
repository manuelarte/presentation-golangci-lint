package main

import (
	"fmt"
	"go/ast"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(NewAnalyzer())
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "astexample",
		Doc:      "astexample to show hierarchy",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch casted := n.(type) {
			case *ast.File:
				fmt.Printf("Package: %q\n", casted.Name.Name)
			case *ast.Ident:
				fmt.Printf("Ident: %q\n", casted.Name)
			case *ast.ImportSpec:
				fmt.Printf("ImportSpec: %s\n", casted.Path.Value)
			case *ast.FuncDecl:
				fmt.Printf("FuncDecl: %q Params(%d) Returns(%d)\n", casted.Name.Name, len(casted.Type.Params.List), len(casted.Type.Results.List))
			default:
				fmt.Printf("Node(%q)\n", reflect.TypeOf(n))
			}
			return true
		})
	}
	return nil, nil
}
