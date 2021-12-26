package nakedliteral

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

var Analyzer = &analysis.Analyzer{
	Name:             "nakedliteral",
	Doc:              "checks for naked literal passed to functions that requries values of the specific defined type",
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		fn, _ := typeutil.Callee(pass.TypesInfo, call).(*types.Func)
		if fn == nil {
			return
		}
		for _, arg := range call.Args {
			typ := pass.TypesInfo.Types[arg].Type

			_, isNamed := typ.(*types.Named)
			_, isUnderlyingBasic := typ.Underlying().(*types.Basic)
			_, isLiteral := arg.(*ast.BasicLit)

			if isNamed && isUnderlyingBasic && isLiteral {
				msg := fmt.Sprintf("passing naked literal to parameter of defined type %q", typ.String())
				pass.Report(analysis.Diagnostic{
					Pos:     arg.Pos(),
					End:     arg.End(),
					Message: msg,
				})
			}
		}
	})
	return nil, nil
}
