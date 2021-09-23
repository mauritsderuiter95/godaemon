package nightmode

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	fmt.Printf("%#v\n", core.State{EntityId: "light.woonkamer"}.Get())
	core.RunDaily(22, 22, n.TurnOffLivingRoom)
}

func (Nightmode) TurnOffLivingRoom() {
	core.Entity{Name: "light.woonkamer"}.TurnOff()
}

func (n Nightmode) ToggleKitchen(event core.Event) {
	core.Entity{Name: "light.kitchen"}.Toggle()
}
