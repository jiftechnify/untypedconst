package callexpr

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

func takesExString(s external.ExString)   {}
func takesExInt(i external.ExInt)         {}
func takesExFloat(f external.ExFloat)     {}
func takesExComplex(c external.ExComplex) {}
func takesExRune(r external.ExRune)       {}
func takesExBool(b external.ExBool)       {}
func takesBasicType(s string, i int, f float64, c complex128, r rune, b bool)

const (
	one = iota + 1
	two
)

const (
	exOne external.ExInt = iota + 1
	exTwo
)

const compl = 0.1 + 0.2i

func Call() {
	takesExString(external.ExString("hoge"))
	takesExString(external.ExString("fug" + "a"))
	takesExString(external.ExStr)
	takesExString("hoge")              // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	takesExString("fug" + "a")         // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`
	takesExString(external.UntypedStr) // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`

	// no warning for untyped const passed as "external & private" defined type
	external.TakesPrivExString("priv")

	takesExInt(external.ExInt((1 + 2) * 3))
	takesExInt(external.ExInt(1) << 10)
	takesExInt(exOne)
	takesExInt(exTwo)
	takesExInt((1 + 2) * 3) // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	takesExInt(1 << 10)     // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	takesExInt(one)         // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`
	takesExInt(two)         // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

	takesExFloat(external.ExFloat(0.1))
	takesExFloat(0.1)         // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	takesExFloat(real(compl)) // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`
	takesExFloat(imag(compl)) // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`

	takesExComplex(external.ExComplex(1.0 + 0.5i))
	takesExComplex(1.0i)              // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	takesExComplex(1.0 + 0.5i)        // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`
	takesExComplex(complex(1.0, 0.5)) // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`

	takesExRune(external.ExRune('a'))
	takesExRune('a')       // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	takesExRune('a' + 1)   // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`
	takesExRune('z' - 'a') // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`

	takesExBool(external.ExBool(true))
	takesExBool(true)                // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	takesExBool(false)               // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	takesExBool(1 > 2)               // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
	takesExBool("fug"+"a" == "fuga") // want `passing untyped constant to parameter of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`

	takesBasicType("a", 1, 1.0, complex(1.0, 2.0), 'a', true)
}
