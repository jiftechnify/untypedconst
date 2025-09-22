package ifstmt

import "github.com/jiftechnify/untypedconst/pkg/external"

func IfStmt(i int, j int, exInt external.ExInt) {
	if i == 0 {
	}
	if exInt == external.ExInt(0) {
	}
	if exInt == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if 0 <= exInt { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}

	if i == 0 && exInt == external.ExInt(0) {
	}
	if exInt == 0 && i == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if i == 0 || exInt == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if i == 0 || (j == 0 && exInt == 0) { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}

func genInt() int {
	return 0
}

func genExInt() external.ExInt {
	return external.ExInt(0)
}

func IfStmtFuncCall() {
	if genInt() == 0 {
	}
	if genExInt() == external.ExInt(0) {
	}
	if genExInt() == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if 0 <= genExInt() { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}

type foo struct {
	i     int
	exInt external.ExInt
}

func (f *foo) getI() int {
	return f.i
}

func (f *foo) getExInt() external.ExInt {
	return f.exInt
}

func IfStmtSelector(f *foo) {
	if f.i == 0 {
	}
	if f.exInt == external.ExInt(0) {
	}
	if f.exInt == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if 0 <= f.exInt { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}

func IfStmtMethodCall(f *foo) {
	if f.getI() == 0 {
	}
	if f.getExInt() == external.ExInt(0) {
	}
	if f.getExInt() == 0 { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	if 0 <= f.getExInt() { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}
