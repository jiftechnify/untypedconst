package returnstmt

import "github.com/jiftechnify/untypedconst/pkg/external"

func retExString() external.ExString {
	if true {
		return external.ExString("hoge")
	} else {
		return "hoge" // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
}

func retExInt() external.ExInt {
	if true {
		return external.ExInt(1)
	} else {
		return 1 // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
}

func retExFloat() external.ExFloat {
	if true {
		return external.ExFloat(0.5)
	} else {
		return 0.5 // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
}

func retExComplex() external.ExComplex {
	if true {
		return external.ExComplex(1.0 + 0.5i)
	} else {
		return 1.0 + 0.5i // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
}

func retExRune() external.ExRune {
	if true {
		return external.ExRune('a')
	} else {
		return 'a' // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
}

func retExBool() external.ExBool {
	if true {
		return external.ExBool(true)
	} else {
		return true // want `returning untyped constant as defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
}
