package main

import (
	"github.com/jiftechnify/nakedliteral"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(nakedliteral.Analyzer)
}
