package types

import (
	"fmt"
	"plugin"
)

/*
Plugin The interface that plugins need to implement so that they can
provide data to moniteur.
*/
type Plugin interface {
	/*
		Schedule Returns the schedule object that includes all the schedule data.
	*/
	Schedule(room string) (*Schedule, error, string)
}

func LoadPlugin(path string) (Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load plugin from %s: %v", path, err)
	}

	pluginSymbol, err := plug.Lookup("PLUGIN")
	if err != nil {
		return nil, fmt.Errorf("failed to load PLUGIN from plugin: %v", err)
	}

	return pluginSymbol.(Plugin), nil
}
