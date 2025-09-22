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
