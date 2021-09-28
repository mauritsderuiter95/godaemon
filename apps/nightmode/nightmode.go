package nightmode

import (
	"fmt"
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"time"
)

type Nightmode struct {
	core.App
}

func (n Nightmode) Initialize() {
	conf, err := n.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	core.RunEveryDay(0, 0, func() {
		day := time.Now().Weekday()
		if day == time.Saturday || day == time.Sunday {
			core.RunIn(1, 0, func() {
				n.TurnOffEntities(conf["entities"].([]string))
			})
		} else {
			n.TurnOffEntities(conf["entities"].([]string))
		}
	})
}

func (n Nightmode) TurnOffEntities(entities []string) {
	for _, entity := range entities {
		core.Entity{EntityId: entity}.TurnOff()
	}
}

func (n Nightmode) ToggleKitchen(core.Event) {
	core.Entity{EntityId: "light.kitchen"}.Toggle()
}
