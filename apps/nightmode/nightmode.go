package nightmode

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

type Nightmode struct {
}

func (n Nightmode) Initialize() {
	core.RunEveryDay(22, 22, n.TurnOffLivingRoom)
	core.Event{EventType: "deconz_event"}.On("switch_woonkamer", n.ToggleKitchen)
}

func (Nightmode) TurnOffLivingRoom() {
	core.Entity{EntityId: "light.woonkamer"}.TurnOff()
}

func (n Nightmode) ToggleKitchen(event core.Event) {
	core.Entity{EntityId: "light.kitchen"}.Toggle()
}
