package switchstmt

type ZeroOne int

const (
	Zero ZeroOne = 0
	One  ZeroOne = 1
)

func SwitchStmt(i int, zeroOne ZeroOne) {
	switch i {
	case 0:
	case 1:
	default:
	}

	switch zeroOne {
	case Zero:
	case ZeroOne(1):
	default:
	}

	switch zeroOne {
	case 0: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	case 1: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	default:
	}
}

func genZeroOne(b bool) ZeroOne {
	if b {
		return One
	}
	return Zero
}

func SwitchStmtFuncCall(b bool) {
	switch genZeroOne(b) {
	case Zero:
	case ZeroOne(1):
	default:
	}

	switch genZeroOne(b) {
	case 0: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	case 1: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	default:
	}
}

type foo struct {
	zeroOne ZeroOne
}

func (f *foo) getZeroOne() ZeroOne {
	return f.zeroOne
}

func SwitchStmtSelector(f *foo) {
	switch f.getZeroOne() {
	case Zero:
	case ZeroOne(1):
	default:
	}

	switch f.getZeroOne() {
	case 0: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	case 1: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	default:
	}
}

func SwitchStmtMethodCall(f *foo) {
	switch f.getZeroOne() {
	case Zero:
	case ZeroOne(1):
	default:
	}

	switch f.getZeroOne() {
	case 0: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	case 1: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	default:
	}
}
