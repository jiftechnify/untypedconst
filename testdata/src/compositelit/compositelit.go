package compositelit

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

var (
	exStringSlice = []external.ExString{
		external.ExString("a"),
		"a", // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	exIntSlice = []external.ExInt{
		external.ExInt(1),
		1, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	exFloatSlice = []external.ExFloat{
		external.ExFloat(0.5),
		0.5, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	exComplexSlice = []external.ExComplex{
		external.ExComplex(1.0 + 0.5i),
		1.0 + 0.5i, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	exRuneSlice = []external.ExRune{
		external.ExRune('a'),
		'a', // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	exBoolSlice = []external.ExBool{
		external.ExBool(true),
		true, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)

var (
	mapExStringKey = map[external.ExString]struct{}{
		external.ExString("a"): {},
		"b":                    {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	mapExIntKey = map[external.ExInt]struct{}{
		external.ExInt(1): {},
		2:                 {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	mapExFloatKey = map[external.ExFloat]struct{}{
		external.ExFloat(0.1): {},
		0.2:                   {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	mapExComplexKey = map[external.ExComplex]struct{}{
		external.ExComplex(1i): {},
		2i:                     {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	mapExRuneKey = map[external.ExRune]struct{}{
		external.ExRune('a'): {},
		'b':                  {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	mapExBoolKey = map[external.ExBool]struct{}{
		external.ExBool(true): {},
		false:                 {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)

var (
	mapExStringVal = map[string]external.ExString{
		"typed":   external.ExString("a"),
		"untyped": "b", // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	mapExIntVal = map[string]external.ExInt{
		"typed":   external.ExInt(1),
		"untyped": 2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	mapExFloatVal = map[string]external.ExFloat{
		"typed":   external.ExFloat(0.1),
		"untyped": 0.2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	mapExComplexVal = map[string]external.ExComplex{
		"typed":   external.ExComplex(1i),
		"untyped": 2i, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	mapExRuneVal = map[string]external.ExRune{
		"typed":   external.ExRune('a'),
		"untyped": 'b', // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	mapExBoolVal = map[string]external.ExBool{
		"typed":   external.ExBool(true),
		"untyped": false, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)
