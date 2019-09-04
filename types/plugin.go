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
	Initialize(examsLink string)
	Schedule() (*Schedule, error)
	ScheduleRoom(room string) (*Schedule, error, string)
	ExamsSchedule() (*Schedule, error)
	ExamsScheduleRoom(room string) (*Schedule, error, string)
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
