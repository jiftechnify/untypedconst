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
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.CallExpr:
			processCallExpr(pass, n)

		case *ast.ReturnStmt:
			processReturnStmt(pass, n)

		case *ast.SendStmt:
			processSendStmt(pass, n)

		case *ast.CompositeLit:
			processCompositeLit(pass, n)
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
		checkAndReport(pass, arg, "passing naked literal to parameter of Defined Type %q")
	}
}

func processReturnStmt(pass *analysis.Pass, ret *ast.ReturnStmt) {
	for _, res := range ret.Results {
		checkAndReport(pass, res, "returning naked literal as Defiend Type %q")
	}
}

func processSendStmt(pass *analysis.Pass, send *ast.SendStmt) {
	checkAndReport(pass, send.Value, "sending naked literal to channel of Defiend Type %q")
}

func processCompositeLit(pass *analysis.Pass, comp *ast.CompositeLit) {
	for _, elt := range comp.Elts {
		switch e := elt.(type) {
		case *ast.KeyValueExpr: // elt is "key: value" form (element of map/struct)
			checkAndReport(pass, e.Key, "using naked literal as composite literal's element key of Defined Type %q")
			checkAndReport(pass, e.Value, "using naked literal as composite literal's element value of Defined Type %q")

		default: // elt is not "key: value" form (element of slice/array)
			checkAndReport(pass, e, "using naked literal as composite literal's element of Defined Type %q")
		}
	}
}

// check if the expression is target of warning and report problems.
func checkAndReport(pass *analysis.Pass, expr ast.Expr, msgfmt string) {
	exprPkgPath := pass.Pkg.Path()
	declType := pass.TypesInfo.Types[expr].Type

	namedTyp, isNamed := declType.(*types.Named)
	if !isNamed {
		return
	}
	if !isBasicOrParenedBasic(expr) {
		return
	}
	if _, isUnderlyingBasic := declType.Underlying().(*types.Basic); !isUnderlyingBasic {
		return
	}
	// For now, (1) the type of expr is declared as Defined Type, (2) the expr is a literal of basic type, and (3) the Underlying Type of declared expr type is basic type

	// expr is target of warning if the declared type of expr is *not* "external package private type"
	if namedTyp.Obj().Exported() || namedTyp.Obj().Pkg().Path() == exprPkgPath {
		pass.Report(analysis.Diagnostic{
			Pos:     expr.Pos(),
			End:     expr.End(),
			Message: fmt.Sprintf(msgfmt, declType.String()),
		})
	}
}

func isBasicOrParenedBasic(expr ast.Expr) bool {
	currExpr := expr
	for {
		switch e := currExpr.(type) {
		case *ast.BasicLit:
			return true

		case *ast.ParenExpr:
			currExpr = e.X
			// continue

		default:
			return false
		}
	}
}
