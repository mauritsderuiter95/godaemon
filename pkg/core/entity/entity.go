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

func (e Entity) OnChange(f func(event core.HaEvent)) {
	ha := core.GetInstance()

	ha.Callbacks[e.Name] = append(ha.Callbacks[e.Name], f)
}

func (e Entity) TurnOn() {
	ha := core.GetInstance()

	if err := ha.CallService("light", "turn_on", e.Name, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) TurnOff() {
	ha := core.GetInstance()

	if err := ha.CallService("light", "turn_off", e.Name, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) Toggle() {
	ha := core.GetInstance()

	if err := ha.CallService("light", "toggle", e.Name, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}
