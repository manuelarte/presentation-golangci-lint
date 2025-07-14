package unexportedconstantscheck

import (
	"fmt"
	"go/ast"

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
		if valueSpec, ok := n.(*ast.ValueSpec); ok {
			// TODO(manuelarte): to be done
			//nolint:forbidigo // remove later
			fmt.Printf("%+v\n", valueSpec)
			// End TODO(manuelarte)
		}
	})

	//nolint:nilnil //any, error
	return nil, nil
}

//nolint:unused // to be used later
func newUnexportedConstantsCheckDiag(i *ast.Ident) analysis.Diagnostic {
	msg := fmt.Sprintf("unexported constant %q should be prefixed with _",
		i.Name)

	return analysis.Diagnostic{
		Pos:     i.Pos(),
		End:     i.End(),
		Message: msg,
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "unexported constant %q should be prefixed with _",
				TextEdits: []analysis.TextEdit{
					{
						Pos: i.Pos(),
						End: i.End(),
						// TODO(manuelarte): Can someone see the problem of fixing it?
						NewText: fmt.Appendf(nil, "_%s", i.Name),
					},
				},
			},
		},
	}
}
