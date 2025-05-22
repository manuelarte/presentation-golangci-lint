package unexportedconstantscheck

import (
	"fmt"
	"go/ast"
	"iter"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "unexportedconstantscheck",
		Doc:      "unexportedconstantscheck checks if unexported constants starts with _",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.ValueSpec)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		// TODO(manuelarte): to be done
	})

	//nolint:nilnil //any, error
	return nil, nil
}

func newUnexportedConstantsCheckDiag(vs *ast.ValueSpec) analysis.Diagnostic {
	msg := fmt.Sprintf("unexported constant %q should be prefixed with _",
		strings.Join(getName(vs.Names), ","))
	return analysis.Diagnostic{
		Pos:            vs.Pos(),
		Message:        msg,
		SuggestedFixes: nil,
	}
}

func getName(is []*ast.Ident) []string {
	var ns = make([]string, len(is))
	for i, ident := range is {
		ns[i] = ident.Name
	}
	return ns
}

func transform[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			if !yield(f(a)) {
				return
			}
		}
	}
}
