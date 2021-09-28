package hooks

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"time"
)

type Hooks struct {
	core.App
}

func (h Hooks) Initialize() {
	conf, err := h.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	for entity, hook := range conf["entities"].(map[string]interface{}) {
		if hook == "dim" {
			core.Entity{EntityId: entity}.AddHook(h.SetNightlyBrightness)
		}
		if hook == "block" {
			core.Entity{EntityId: entity}.AddHook(h.BlockNightlyTurnOn)
		}
	}
}

func (h Hooks) BlockNightlyTurnOn(e core.Entity) core.State {
	newState := core.State{EntityId: e.EntityId}

	t := time.Now()
	if t.Hour() > 20 || t.Hour() < 6 {
		newState.Attributes = map[string]interface{}{
			"brightness": 0,
		}
	}

	return newState
}

func (h Hooks) SetNightlyBrightness(e core.Entity) core.State {
	newState := core.State{EntityId: e.EntityId}

	t := time.Now()
	if t.Hour() > 20 || t.Hour() < 6 {
		newState.Attributes = map[string]interface{}{
			"brightness": 1,
		}
	} else {
		newState.Attributes = map[string]interface{}{
			"brightness": 140,
		}
	}

	return newState
}
