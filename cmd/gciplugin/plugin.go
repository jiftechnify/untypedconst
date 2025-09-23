package gciplugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/jiftechnify/untypedconst"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("untypedconst", New)
}

func New(conf any) (register.LinterPlugin, error) {
	return &Plugin{}, nil
}

type Plugin struct{}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{untypedconst.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
