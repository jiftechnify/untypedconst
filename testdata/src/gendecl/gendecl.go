package gendecl

import (
	"fmt"

	"github.com/jiftechnify/untypedconst/pkg/external"
)

var ExStringOK = external.ExString("a")
var ExStringNG external.ExString = "a" // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`

var (
	ExIntOK                = external.ExInt(1)
	ExIntNG external.ExInt = 1 // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

	ExFloatOK                  = external.ExFloat(0.5)
	ExFloatNG external.ExFloat = 0.5 // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`

	ExComplexOK                    = external.ExComplex(1.0 + 0.5i)
	ExComplexNG external.ExComplex = 1.0 + 0.5i // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`

	ExRuneOK                 = external.ExRune('a')
	ExRuneNG external.ExRune = 'a' // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`

	ExBoolOK                 = external.ExBool(true)
	ExBoolNG external.ExBool = true // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
)

func LocalVar() {
	var lExStringOK = external.ExString("a")
	var lExStringNG external.ExString = "a" // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	var (
		lExIntOK                = external.ExInt(1)
		lExIntNG external.ExInt = 1 // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

		lExFloatOK                  = external.ExFloat(0.5)
		lExFloatNG external.ExFloat = 0.5 // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`

		lExComplexOK                    = external.ExComplex(1.0 + 0.5i)
		lExComplexNG external.ExComplex = 1.0 + 0.5i // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`

		lExRuneOK                 = external.ExRune('a')
		lExRuneNG external.ExRune = 'a' // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`

		lExBoolOK                 = external.ExBool(true)
		lExBoolNG external.ExBool = true // want `assigning untyped constant to variable of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	)
	// suppress "unused" warnings
	fmt.Println(lExStringOK, lExStringNG, lExIntOK, lExIntNG, lExFloatOK, lExFloatNG, lExComplexOK, lExComplexNG, lExRuneOK, lExRuneNG, lExBoolOK, lExBoolNG)
}
