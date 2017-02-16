package timeline

import (
	"encoding/json"
	"os"
	"os/signal"
	"testing"

	"../../model"
)

func TestTimeline(t *testing.T) {
	channel := make(chan model.Task, 1)
	ts := []model.Task{
		{Option: model.TaskOption{Color: "1"}, Type: "sound", Start: "1"},
		{Option: model.TaskOption{Color: "2"}, Type: "sound", Start: "2"},
		{Option: model.TaskOption{Color: "3"}, Type: "sound", Start: "3"},
	}
	Start(channel, ts)

	signalChannel := make(chan os.Signal, 10)
	signal.Notify(signalChannel, os.Interrupt)
	count := 0
	for {
		select {
		case b := <-channel:
			json.Marshal(b)
			count++
			if count == len(ts) {
				os.Exit(0)
			}
		case <-signalChannel:
			return
		}
	}
}
