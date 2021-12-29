package main

import (
	"github.com/jiftechnify/untypedconst"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(untypedconst.Analyzer)
}
