package switchstmt

type ZeroOne int

const (
	Zero ZeroOne = 0
	One  ZeroOne = 1
)

func SwitchStmt(i int, zo ZeroOne) {
	switch i {
	case 0:
	case 1:
	default:
	}

	switch zo {
	case Zero:
	case ZeroOne(1):
	default:
	}

	switch zo {
	case 0: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	case 1: // want `matching untyped constant with expression of defined type "switchstmt.ZeroOne"`
	default:
	}
}
