package nightmode

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"github.com/mauritsderuiter95/godaemon/pkg/core/entity"
	"github.com/mauritsderuiter95/godaemon/pkg/core/event"
	"github.com/mauritsderuiter95/godaemon/pkg/core/schedule"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	event.Get("deconz_event").OnChange("switch_woonkamer", n.ToggleKitchen)
	schedule.RunDaily(22, 22, n.TurnOffLivingRoom)
}

func (Nightmode) TurnOffLivingRoom() {
	entity.Get("light.woonkamer").TurnOff()
}

func (n Nightmode) ToggleKitchen(event core.HaEvent) {
	entity.Get("light.kitchen").Toggle()
}
