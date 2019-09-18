package main

import (
	"github.com/andreiubr/yab/plugin"
	flags "github.com/jessevdk/go-flags"
)

// provides a bridge between the plugin::Parser interface and the go-flags::Parser struct
type pluginParserAdapter struct {
	*flags.Parser
}

func (p pluginParserAdapter) AddFlagGroup(shortDescription, longDescription string, data interface{}) error {
	_, err := p.AddGroup(shortDescription, longDescription, data)
	return err
}

var _ plugin.Parser = (*pluginParserAdapter)(nil)
