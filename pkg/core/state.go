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

func (s State) Get() State {
	ha := GetInstance()

	return ha.GetState(s.EntityId)
}

func (s State) OnChange(f func(event Event)) {
	ha := GetInstance()

	ha.Callbacks[s.EntityId] = append(ha.Callbacks[s.EntityId], func(event Event) {
		if event.EventType == s.EntityId {
			f(event)
		}
	})
}
