package indexexpr

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

var (
	mapExStringKey = map[external.ExString]struct{}{
		external.ExString("a"): {},
	}
	mapExIntKey = map[external.ExInt]struct{}{
		external.ExInt(1): {},
	}
	mapExFloatKey = map[external.ExFloat]struct{}{
		external.ExFloat(0.1): {},
	}
	mapExComplexKey = map[external.ExComplex]struct{}{
		external.ExComplex(1i): {},
	}
	mapExRuneKey = map[external.ExRune]struct{}{
		external.ExRune('a'): {},
	}
	mapExBoolKey = map[external.ExBool]struct{}{
		external.ExBool(true): {},
	}
)

func Index() {
	_ = mapExStringKey[external.ExString("a")]
	_ = mapExStringKey["a"] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`

	_ = mapExIntKey[external.ExInt(1)]
	_ = mapExIntKey[1] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

	_ = mapExFloatKey[external.ExFloat(0.1)]
	_ = mapExFloatKey[0.1] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`

	_ = mapExComplexKey[external.ExComplex(1i)]
	_ = mapExComplexKey[1i] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`

	_ = mapExRuneKey[external.ExRune('a')]
	_ = mapExRuneKey['a'] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`

	_ = mapExBoolKey[external.ExBool(true)]
	_ = mapExBoolKey[true] // want `using untyped constant for indexing the value whose key type is defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
}
