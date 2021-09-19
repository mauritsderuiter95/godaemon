package nightmode

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"github.com/mauritsderuiter95/godaemon/pkg/core/entity"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	entity.Get("light.woonkamer").TurnOn()
	entity.Get("light.woonkamer").TurnOff()
}
