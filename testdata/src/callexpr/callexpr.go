package callexpr

import "github.com/jiftechnify/nakedliteral/testdata/external"

func takesExString(s external.ExString)   {}
func takesExInt(i external.ExInt)         {}
func takesExFloat(f external.ExFloat)     {}
func takesExComplex(c external.ExComplex) {}
func takesExBool(b external.ExBool)       {}
func takesExRune(r external.ExRune)       {}

func Call() {
	takesExString(external.ExString("hoge"))
	takesExString(external.ExStr)
	takesExString("a")       // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExString"`
	takesExString(("paren")) // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExString"`

	takesExInt(external.ExInt(1))
	takesExInt((1 + 2) + 3) // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExInt"`

	takesExFloat(external.ExFloat(0.1))
	takesExFloat(0.1) // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExFloat"`

	takesExComplex(external.ExComplex(2.0i))
	takesExComplex(2.0i) // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExComplex"`

	takesExBool(external.ExBool(true))
	takesExBool(true) // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExBool"`

	takesExRune(external.ExRune('a'))
	takesExRune('a') // want `passing naked literal to parameter of Defined Type "github.com/jiftechnify/nakedliteral/testdata/external.ExRune"`
}
