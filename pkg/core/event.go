package core

import (
	"time"
)

type Event struct {
	EventType string `json:"event_type"`
	Data      struct {
		Id       string `json:"id"`
		EntityId string `json:"entity_id"`
		OldState struct {
			EntityId    string                 `json:"entity_id"`
			State       string                 `json:"state"`
			Attributes  map[string]interface{} `json:"attributes"`
			LastChanged time.Time              `json:"last_changed"`
			LastUpdated time.Time              `json:"last_updated"`
			Context     struct {
				Id       string      `json:"id"`
				ParentId interface{} `json:"parent_id"`
				UserId   string      `json:"user_id"`
			} `json:"context"`
		} `json:"old_state"`
		NewState struct {
			EntityId    string                 `json:"entity_id"`
			State       string                 `json:"state"`
			Attributes  map[string]interface{} `json:"attributes"`
			LastChanged time.Time              `json:"last_changed"`
			LastUpdated time.Time              `json:"last_updated"`
			Context     struct {
				Id       string      `json:"id"`
				ParentId interface{} `json:"parent_id"`
				UserId   string      `json:"user_id"`
			} `json:"context"`
		} `json:"new_state"`
		Domain      string `json:"domain"`
		Service     string `json:"service"`
		ServiceData struct {
			EntityId interface{} `json:"entity_id"`
		} `json:"service_data"`
	} `json:"data"`
	Origin    string    `json:"origin"`
	TimeFired time.Time `json:"time_fired"`
	Context   struct {
		Id       string      `json:"id"`
		ParentId interface{} `json:"parent_id"`
		UserId   string      `json:"user_id"`
	} `json:"context"`
}

func (e Event) OnTrigger(f func(event Event)) {
	ha := GetInstance()

	if _, ok := ha.Callbacks[e.EventType]; !ok {
		ha.Callbacks[e.EventType] = map[string][]func(Event){}
	}

	ha.Callbacks[e.EventType]["all"] = append(ha.Callbacks[e.EventType]["all"], func(event Event) {
		if event.EventType == e.EventType {
			f(event)
		}
	})
}

func (e Event) OnEntityTrigger(entity string, f func(event Event)) {
	ha := GetInstance()

	if _, ok := ha.Callbacks[e.EventType]; !ok {
		ha.Callbacks[e.EventType] = map[string][]func(Event){}
	}

	ha.Callbacks[e.EventType][entity] = append(ha.Callbacks[e.EventType][entity], func(event Event) {
		if event.EventType == e.EventType {
			f(event)
		}
	})
}
