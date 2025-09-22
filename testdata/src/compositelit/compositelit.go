package compositelit

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

var (
	sliceExString = []external.ExString{
		external.ExString("a"),
		"a", // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	sliceExInt = []external.ExInt{
		external.ExInt(1),
		1, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	sliceExFloat = []external.ExFloat{
		external.ExFloat(0.5),
		0.5, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	sliceExComplex = []external.ExComplex{
		external.ExComplex(1.0 + 0.5i),
		1.0 + 0.5i, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	sliceExRune = []external.ExRune{
		external.ExRune('a'),
		'a', // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	sliceExBool = []external.ExBool{
		external.ExBool(true),
		true, // want `using untyped constant as composite literal's element of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)

var (
	mapKeyExString = map[external.ExString]struct{}{
		external.ExString("a"): {},
		"b":                    {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	mapKeyExInt = map[external.ExInt]struct{}{
		external.ExInt(1): {},
		2:                 {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	mapKeyExFloat = map[external.ExFloat]struct{}{
		external.ExFloat(0.1): {},
		0.2:                   {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	mapKeyExComplex = map[external.ExComplex]struct{}{
		external.ExComplex(1i): {},
		2i:                     {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	mapKeyExRune = map[external.ExRune]struct{}{
		external.ExRune('a'): {},
		'b':                  {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	mapKeyExBool = map[external.ExBool]struct{}{
		external.ExBool(true): {},
		false:                 {}, // want `using untyped constant as composite literal's element key of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)

var (
	mapValExString = map[string]external.ExString{
		"typed":   external.ExString("a"),
		"untyped": "b", // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	mapValExInt = map[string]external.ExInt{
		"typed":   external.ExInt(1),
		"untyped": 2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	mapValExFloat = map[string]external.ExFloat{
		"typed":   external.ExFloat(0.1),
		"untyped": 0.2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	mapValExComplex = map[string]external.ExComplex{
		"typed":   external.ExComplex(1i),
		"untyped": 2i, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	mapValExRune = map[string]external.ExRune{
		"typed":   external.ExRune('a'),
		"untyped": 'b', // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	mapValExBool = map[string]external.ExBool{
		"typed":   external.ExBool(true),
		"untyped": false, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)

type GenericStruct[T any] struct {
	Typed   T
	Untyped T
}

var (
	genStructExString = GenericStruct[external.ExString]{
		Typed:   external.ExString("a"),
		Untyped: "b", // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	}
	genStructExInt = GenericStruct[external.ExInt]{
		Typed:   external.ExInt(1),
		Untyped: 2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	}
	genStructExFloat = GenericStruct[external.ExFloat]{
		Typed:   external.ExFloat(0.1),
		Untyped: 0.2, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	}
	genStructExComplex = GenericStruct[external.ExComplex]{
		Typed:   external.ExComplex(1i),
		Untyped: 2i, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	}
	genStructExRune = GenericStruct[external.ExRune]{
		Typed:   external.ExRune('a'),
		Untyped: 'b', // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	}
	genStructExBool = GenericStruct[external.ExBool]{
		Typed:   external.ExBool(true),
		Untyped: false, // want `using untyped constant as composite literal's element value of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	}
)
