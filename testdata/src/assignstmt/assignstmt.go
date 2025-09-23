package assignstmt

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

func AssignStmt() {
	var (
		i     int
		exInt external.ExInt
	)

	i = 42
	exInt = external.ExInt(42)
	exInt = 42 // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

	// suppress unused warnings
	print(i, exInt)
}

func AssignStmtInFor() {
	var (
		i     int
		exInt external.ExInt
	)

	for i = 0; i < 10; i++ {
	}
	for exInt = external.ExInt(0); exInt < external.ExInt(10); exInt++ {
	}
	for exInt = 0; exInt < 10; exInt++ { // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}
