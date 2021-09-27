package nightmode

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

type Nightmode struct {
}

func (n Nightmode) Initialize() {
	state, err := core.State{EntityId: "light.woonkamer"}.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", state)
	core.RunEveryDay(22, 22, n.TurnOffLivingRoom)
	core.Event{EventType: "deconz_event"}.On("switch_woonkamer", n.ToggleKitchen)
}

func (Nightmode) TurnOffLivingRoom() {
	core.Entity{EntityId: "light.woonkamer"}.TurnOff()
}

func (n Nightmode) ToggleKitchen(event core.Event) {
	core.Entity{EntityId: "light.kitchen"}.Toggle()
}
