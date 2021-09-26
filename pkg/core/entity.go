package core

import (
	"log"
)

type Entity struct {
	Name string
}

func (e Entity) OnChange(f func(event Event)) {
	ha := GetInstance()

	ha.Callbacks[e.Name] = append(ha.Callbacks[e.Name], f)
}

func (e Entity) TurnOn(attrs map[string]string) {
	ha := GetInstance()

	if err := ha.CallService("light", "turn_on", e.Name, attrs); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) TurnOff() {
	ha := GetInstance()

	if err := ha.CallService("light", "turn_off", e.Name, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) Toggle() {
	ha := GetInstance()

	if err := ha.CallService("light", "toggle", e.Name, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}
