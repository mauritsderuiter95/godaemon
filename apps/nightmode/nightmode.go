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
				n.TurnOffEntities(conf["entities"].([]interface{}))
			})
		} else {
			n.TurnOffEntities(conf["entities"].([]interface{}))
		}
	})
}

func (n Nightmode) TurnOffEntities(entities []interface{}) {
	for _, entity := range entities {
		core.Entity{EntityId: entity.(string)}.TurnOff()
	}
}

func (n Nightmode) ToggleKitchen(core.Event) {
	core.Entity{EntityId: "light.kitchen"}.Toggle()
}
