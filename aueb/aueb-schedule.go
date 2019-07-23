package main

import (
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
)

// PLUGIN The plugin to be read by the moniteur agent.
var PLUGIN = Plugin{}

type Plugin struct {
}

func (Plugin) Schedule() (*types.Schedule, error) {
	return nil, fmt.Errorf("this must be working")
}