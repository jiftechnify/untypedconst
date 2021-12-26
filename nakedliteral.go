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
	// path of analyzing package
	pkgPath := pass.Pkg.Path()

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

			if argIsLintTarget(pkgPath, arg, typ) {
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

// check if the function argument is target of warning
func argIsLintTarget(callerPkgPath string, argExpr ast.Expr, argType types.Type) bool {
	namedTyp, isNamed := argType.(*types.Named)
	if !isNamed {
		return false
	}
	if _, isLiteral := argExpr.(*ast.BasicLit); !isLiteral {
		return false
	}
	if _, isUnderlyingBasic := argType.Underlying().(*types.Basic); !isUnderlyingBasic {
		return false
	}
	// For now, (1) the type of func arg is Defined Type, (2) the arg is a literal of basic type, and (3) the Underlying Type of arg type is basic type

	// arg is target of warning if the type of arg is *not* "external package private type"
	return namedTyp.Obj().Exported() || namedTyp.Obj().Pkg().Path() == callerPkgPath
}
