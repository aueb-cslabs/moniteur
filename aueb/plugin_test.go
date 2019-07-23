package main

import (
	"fmt"
	"testing"
)

func TestPlugin_Schedule(t *testing.T) {
	schedule, _ := PLUGIN.Schedule()
	fmt.Println(len(schedule.Slots))
}
