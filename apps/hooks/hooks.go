package hooks

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"time"
)

type Hooks struct{}

func (h Hooks) Initialize() {
	core.Entity{EntityId: "light.bulb_overloop"}.AddHook(h.BlockNightlyTurnOn)
}

func (h Hooks) BlockNightlyTurnOn() core.State {
	newState := core.State{EntityId: "light.bulb_overloop"}

	t := time.Now()
	if t.Hour() > 20 || t.Hour() < 6 {
		newState.Attributes = map[string]interface{}{
			"brightness": 0,
		}
	}

	return newState
}
