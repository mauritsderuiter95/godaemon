package entity

import (
	"github.com/mauritsderuiter95/godaemon/pkg/core"
	"log"
)

type Entity struct {
	Name string
}

func Get(name string) Entity {
	return Entity{Name: name}
}

func (e Entity) TurnOn() {
	ha := core.GetInstance()

	if err := ha.SendMessage("call_service", "light", "turn_on", "", "", e.Name); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}
