package nightmode

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"github.com/mauritsderuiter95/godaemon/pkg/core/entity"
	"github.com/mauritsderuiter95/godaemon/pkg/core/event"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	event.Get("deconz_event").OnChange("switch_woonkamer", n.ToggleKitchen)
}

func (n Nightmode) ToggleKitchen(event core.HaEvent) {
	entity.Get("light.kitchen").Toggle()
}
