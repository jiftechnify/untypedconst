package untypedconst_test

import (
	"testing"

	"github.com/jiftechnify/untypedconst"
	"golang.org/x/tools/go/analysis/analysistest"
)

var pkgs = []string{
	"github.com/...",
	"callexpr",
	"sendstmt",
	"returnstmt",
	"compositelit",
	"indexexpr",
	"gendecl",
	"ifstmt",
	"switchstmt",
}

func TestAll(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, untypedconst.Analyzer, pkgs...)
}
