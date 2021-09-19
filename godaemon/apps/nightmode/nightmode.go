package nightmode

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"github.com/mauritsderuiter95/godaemon/pkg/core/entity"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	entity.Get("light.woonkamer").TurnOn()
	entity.Get("light.woonkamer").TurnOff()
	entity.Get("light.woonkamer").OnChange(n.SyncKitchen)
}

func (n Nightmode) SyncKitchen(event core.Event) {
	fmt.Println(event.Event.Data.NewState.State)
	if event.Event.Data.NewState.State == "on" {
		entity.Get("light.kitchen").TurnOn()
	} else {
		entity.Get("light.kitchen").TurnOff()
	}
}
