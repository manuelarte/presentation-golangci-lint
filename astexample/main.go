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
			fmt.Printf("Node(%q)\n", reflect.TypeOf(n))
			return true
		})
	}
	return nil, nil
}
