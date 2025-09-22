package untypedconst

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

var Analyzer = &analysis.Analyzer{
	Name:     "untypedconst",
	Doc:      "Checks that untyped constant expressions are used as values of a defined type",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
		(*ast.ReturnStmt)(nil),
		(*ast.SendStmt)(nil),
		(*ast.CompositeLit)(nil),
		(*ast.IndexExpr)(nil),
		(*ast.GenDecl)(nil),
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

		case *ast.IndexExpr:
			processIndexExpr(pass, n)

		case *ast.GenDecl:
			processGenDecl(pass, n)
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
		checkAndReport(pass, arg, "passing untyped constant to parameter of defined type %q")
	}
}

func processReturnStmt(pass *analysis.Pass, ret *ast.ReturnStmt) {
	for _, res := range ret.Results {
		checkAndReport(pass, res, "returning untyped constant as defined type %q")
	}
}

func processSendStmt(pass *analysis.Pass, send *ast.SendStmt) {
	checkAndReport(pass, send.Value, "sending untyped constant to channel of defined type %q")
}

func processCompositeLit(pass *analysis.Pass, comp *ast.CompositeLit) {
	for _, elt := range comp.Elts {
		switch e := elt.(type) {
		case *ast.KeyValueExpr: // elt is "key: value" form (element of map/struct)
			checkAndReport(pass, e.Key, "using untyped constant as composite literal's element key of defined type %q")
			checkAndReport(pass, e.Value, "using untyped constant as composite literal's element value of defined type %q")

		default: // elt is not "key: value" form (element of slice/array)
			checkAndReport(pass, e, "using untyped constant as composite literal's element of defined type %q")
		}
	}
}

func processIndexExpr(pass *analysis.Pass, idx *ast.IndexExpr) {
	checkAndReport(pass, idx.Index, "using untyped constant for indexing the value whose key type is defined type %q")
}

func processGenDecl(pass *analysis.Pass, decl *ast.GenDecl) {
	// check `var` only
	if decl.Tok != token.VAR {
		return
	}
	for _, spec := range decl.Specs {
		valSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}
		for _, val := range valSpec.Values {
			checkAndReport(pass, val, "assigning untyped constant to variable of defined type %q")
		}
	}
}

// check if the expression is target of warning, and report problems.
//
// `msgfmt` MUST contain exact one format specifier for string(`%s` or `%q`)
func checkAndReport(pass *analysis.Pass, expr ast.Expr, msgfmt string) {
	// check if the expr is an untyped constant expression.
	if pass.TypesInfo.Types[expr].Value == nil {
		return
	}
	if !isUntypedConstExpr(pass, expr) {
		return
	}

	// check if the inferred type of the expr is a defined type and its underlying type is basic type.
	inferredType := pass.TypesInfo.Types[expr].Type
	namedTyp, isNamed := inferredType.(*types.Named)
	if !isNamed {
		return
	}
	if _, isUnderlyingBasic := inferredType.Underlying().(*types.Basic); !isUnderlyingBasic {
		return
	}

	// no problem if the inferred type of the expr is "external private type", so exclude such cases.
	if namedTyp.Obj().Exported() || namedTyp.Obj().Pkg().Path() == pass.Pkg.Path() {
		pass.ReportRangef(expr, msgfmt, inferredType.String())
	}
}

// check if `expr` is untyped.
// precondition: `expr` is constant expression (i.e. has constant value).
func isUntypedConstExpr(pass *analysis.Pass, expr ast.Expr) bool {
	switch e := ast.Unparen(expr).(type) {
	case *ast.BasicLit:
		// Naked basic literals are untyped const expr.
		return true

	case *ast.Ident:
		// `true`, `false` and `iota` are untyped const expr.
		if _, isConst := constIdentNames[e.Name]; isConst {
			return true
		}
		// For other identifiers, check if it stands for an untyped constant.
		return isUntypedConstIdent(pass, e)

	case *ast.SelectorExpr:
		// It should be qualified identifier of constant in external package.
		// Check if it stands for is untyped constant.
		return isUntypedConstIdent(pass, e.Sel)

	case *ast.UnaryExpr:
		// If an operand is untyped, entire expression is also untyped.
		return isUntypedConstExpr(pass, e.X)

	case *ast.BinaryExpr:
		// "A constant comparison always yields an untyped boolean constant", as stated in the lang spec.
		if _, isComparison := comparisonTokens[e.Op.String()]; isComparison {
			return true
		}
		// `expr` is other than comparison. In this case, if both operands are untyped, entire expression is also untyped.
		return isUntypedConstExpr(pass, e.X) && isUntypedConstExpr(pass, e.Y)

	case *ast.CallExpr:
		// As stated in the lang spec:
		// * "Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant".
		// * "For `real` and `imag`, ... If the argument evaluates to an untyped constant, it must be a number, and the return value of the function is an untyped floating-point constant".
		// All other call expressions (incl. type conversions) are typed. (any counter examples?)
		callee, ok := typeutil.Callee(pass.TypesInfo, e).(*types.Builtin)
		if !ok {
			return false
		}
		if _, isBuiltInAboutComplex := builtInsAboutComplex[callee.Name()]; !isBuiltInAboutComplex {
			return false
		}

		// Callee is a built-in function one of `complex`, `real`, `imag`.
		for _, arg := range e.Args {
			if !isUntypedConstExpr(pass, arg) {
				return false
			}
		}
		// All args are untyped!
		return true

	default:
		// All other types of expression (index, key-value, slice, star) can't appear in const expr.
		log.Printf("unexpected node type: %T (node: %+v)", e, e)
		return false
	}
}

// Check if `id` stands for an untyped constant.
func isUntypedConstIdent(pass *analysis.Pass, id *ast.Ident) bool {
	cnst, ok := pass.TypesInfo.ObjectOf(id).(*types.Const)
	if !ok {
		return false
	}
	return strings.HasPrefix(cnst.Type().String(), "untyped")
}

var constIdentNames = map[string]struct{}{
	"true":  {},
	"false": {},
	"iota":  {},
}

var comparisonTokens = map[string]struct{}{
	"==": {},
	"!=": {},
	"<":  {},
	"<=": {},
	">":  {},
	">=": {},
}

var builtInsAboutComplex = map[string]struct{}{
	"complex": {},
	"real":    {},
	"imag":    {},
}
