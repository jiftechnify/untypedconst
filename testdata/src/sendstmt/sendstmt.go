package sendstmt

import (
	"github.com/jiftechnify/untypedconst/pkg/external"
)

var (
	exStringChan  = make(chan external.ExString)
	exIntChan     = make(chan external.ExInt)
	exFloatChan   = make(chan external.ExFloat)
	exComplexChan = make(chan external.ExComplex)
	exRuneChan    = make(chan external.ExRune)
	exBoolChan    = make(chan external.ExBool)
)

func Send() {
	exStringChan <- external.ExString("hoge")
	exStringChan <- "hoge" // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExString"`

	// no warning for untyped const sent as "external & private" defined type
	external.PrivExStringChan <- external.PrivExStr

	exIntChan <- external.ExInt(1)
	exIntChan <- 1 // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExInt"`

	exFloatChan <- external.ExFloat(0.5)
	exFloatChan <- 0.5 // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExFloat"`

	exComplexChan <- external.ExComplex(1.0 + 0.5i)
	exComplexChan <- 1.0 + 0.5i // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExComplex"`

	exRuneChan <- external.ExRune('a')
	exRuneChan <- 'a' // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExRune"`

	exBoolChan <- external.ExBool(true)
	exBoolChan <- true // want `sending untyped constant to channel of defined type "github.com/jiftechnify/untypedconst/pkg/external.ExBool"`
}
