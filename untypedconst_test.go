package untypedconst_test

import (
	"testing"

	"github.com/jiftechnify/untypedconst"
	"golang.org/x/tools/go/analysis/analysistest"
)

var pkgs = []string{
	"github.com/...",
	"assignstmt",
	"callexpr",
	"compositelit",
	"gendecl",
	"ifstmt",
	"indexexpr",
	"returnstmt",
	"sendstmt",
	"switchstmt",
}

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, untypedconst.Analyzer, pkgs...)
}
