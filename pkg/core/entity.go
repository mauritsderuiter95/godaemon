package core

import (
	"log"
)

type Entity struct {
	EntityId string
	State    State
}

func GetEntity(entityId string) (Entity, error) {
	state, err := State{EntityId: entityId}.Get()
	if err != nil {
		return Entity{}, err
	}
	return Entity{
		EntityId: entityId,
		State:    state,
	}, nil
}

func (e Entity) OnChange(f func(event Event)) {
	ha := GetInstance()

	if _, ok := ha.Callbacks["all"]; !ok {
		ha.Callbacks["all"] = map[string][]func(Event){}
	}

	ha.Callbacks["all"][e.EntityId] = append(ha.Callbacks["all"][e.EntityId], f)
}

func (e Entity) TurnOn(attrs map[string]interface{}) {
	ha := GetInstance()

	if err := ha.CallService("light", "turn_on", e.EntityId, attrs); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) TurnOff() {
	ha := GetInstance()

	if err := ha.CallService("light", "turn_off", e.EntityId, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) Toggle() {
	ha := GetInstance()

	if err := ha.CallService("light", "toggle", e.EntityId, nil); err != nil {
		logger := log.Default()
		logger.Println(err)
	}
}

func (e Entity) AddHook(f func(e Entity) State) {
	ha := GetInstance()

	ha.Hooks[e.EntityId] = append(ha.Hooks[e.EntityId], func() State {
		return f(e)
	})
}
