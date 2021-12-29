package untypedconst_test

import (
	"testing"

	"github.com/jiftechnify/untypedconst"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, untypedconst.Analyzer, "github.com/...", "callexpr", "sendstmt", "returnstmt", "compositelit", "indexexpr")
}
