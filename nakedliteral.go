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
		(*ast.ReturnStmt)(nil),
		(*ast.SendStmt)(nil),
	}
	inspect.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.CallExpr:
			processCallExpr(pass, n)

		case *ast.ReturnStmt:
			processReturnStmt(pass, n)

		case *ast.SendStmt:
			processSendStmt(pass, n)
		}
	})

	return nil, nil
}

func processCallExpr(pass *analysis.Pass, call *ast.CallExpr) {
	fn, _ := typeutil.Callee(pass.TypesInfo, call).(*types.Func)
	if fn == nil {
		return
	}

	for _, arg := range call.Args {
		typ := pass.TypesInfo.Types[arg].Type

		if exprIsLintTarget(pass.Pkg.Path(), arg, typ) {
			msg := fmt.Sprintf("passing naked literal to parameter of Defined Type %q", typ.String())
			pass.Report(analysis.Diagnostic{
				Pos:     arg.Pos(),
				End:     arg.End(),
				Message: msg,
			})
		}
	}
}

func processReturnStmt(pass *analysis.Pass, ret *ast.ReturnStmt) {
	for _, res := range ret.Results {
		typ := pass.TypesInfo.Types[res].Type

		if exprIsLintTarget(pass.Pkg.Path(), res, typ) {
			msg := fmt.Sprintf("returning naked literal as Defiend Type %q", typ.String())
			pass.Report(analysis.Diagnostic{
				Pos:     res.Pos(),
				End:     res.End(),
				Message: msg,
			})
		}
	}
}

func processSendStmt(pass *analysis.Pass, send *ast.SendStmt) {
	typ := pass.TypesInfo.Types[send.Value].Type

	if exprIsLintTarget(pass.Pkg.Path(), send.Value, typ) {
		msg := fmt.Sprintf("sending naked literal to channel of Defiend Type %q", typ.String())
		pass.Report(analysis.Diagnostic{
			Pos:     send.Value.Pos(),
			End:     send.Value.End(),
			Message: msg,
		})
	}
}

// check if the expression is target of warning
func exprIsLintTarget(callerPkgPath string, expr ast.Expr, declType types.Type) bool {
	namedTyp, isNamed := declType.(*types.Named)
	if !isNamed {
		return false
	}
	if _, isLiteral := expr.(*ast.BasicLit); !isLiteral {
		return false
	}
	if _, isUnderlyingBasic := declType.Underlying().(*types.Basic); !isUnderlyingBasic {
		return false
	}
	// For now, (1) the type of expr is declared as Defined Type, (2) the expr is a literal of basic type, and (3) the Underlying Type of declared expr type is basic type

	// arg is target of warning if the type of arg is *not* "external package private type"
	return namedTyp.Obj().Exported() || namedTyp.Obj().Pkg().Path() == callerPkgPath
}
