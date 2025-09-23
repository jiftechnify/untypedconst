package forstmt

import "github.com/jiftechnify/untypedconst/pkg/external"

func ForStmtCond(i int, exInt external.ExInt) {
	for ; i < 10; i++ {
	}
	for ; exInt < external.ExInt(10); exInt++ {
	}
	for ; exInt < 10; exInt++ { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}

	for ; i < 10 && exInt < external.ExInt(10); i++ {
	}
	for ; i < 10 && exInt < 10; i++ { // want `comparing untyped constant to expression of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}
