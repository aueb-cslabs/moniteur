package main

import (
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
)

func main() {

	config, _ := types.LoadConfiguration("config.yml")
	plugin, _ := types.LoadPlugin(config.Plugin)

	_, err := plugin.Schedule()
	fmt.Println(err)

}
