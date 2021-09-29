package core

import (
	"time"
)

type State struct {
	Attributes  map[string]interface{} `json:"attributes"`
	EntityId    string                 `json:"entity_id"`
	LastChanged time.Time              `json:"last_changed"`
	LastUpdated time.Time              `json:"last_updated"`
	State       string                 `json:"state"`
}

func (s State) Get() (State, error) {
	ha := GetInstance()

	return ha.GetState(s.EntityId)
}

func (s State) OnChange(f func(event Event)) {
	ha := GetInstance()

	if _, ok := ha.Callbacks["all"]; !ok {
		ha.Callbacks["all"] = map[string][]func(Event){}
	}

	ha.Callbacks["all"][s.EntityId] = append(ha.Callbacks["all"][s.EntityId], func(event Event) {
		if event.EventType == s.EntityId {
			f(event)
		}
	})
}
